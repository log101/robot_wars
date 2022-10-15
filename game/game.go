// Uygulamanın interaktif komponentlerini bir araya getirildiği yer
// Oyunun dünyası
package game

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/robot_wars/game/robot"
	"github.com/robot_wars/game/robot_league"
	"github.com/robot_wars/utils"
)

type GameState string

// Oyunun bu alttaki 5 durumdan birinde olmak zorunda
// Her durumda oyun farklı seçenekler sunuyor
var (
	stateWaiting           GameState = "stateWaiting"           // Beklemede
	statePlayingMatch      GameState = "statePlayingMatch"      // Lig başladı, sırasıyla maçlar yapılacak
	stateMatchConfirmation GameState = "stateMatchConfirmation" // Lige başlamak için onay bekleniyor
	stateMatchEnded        GameState = "stateMatchEnded"        // Maç bitti, yeni maça geçilecek
	stateGameOver          GameState = "stateGameOver"          // Oyun bitti, yapılacak maç kalmadı (program kapanır)
)

// Oyun yapısı içerisinde anlık durumu ve robotların
// olduğu ligi tutuyor
type Game struct {
	League *robot_league.RobotLeague
	state  GameState
}

// METODLAR
// Start metodu oyunu başlatıyor, anlık duruma göre switch
// doğrultusunda hareket ediyor
func (g *Game) Start() {
	g.SetState(stateWaiting)       // Başlangıç durumu bekleme
	displayHelpMessage.operation() // Yardım mesajını görüntüle

	for {
		switch g.GetState() {
		case stateWaiting:
			// Beklemedeki seçenekler:
			// - yardım mesajı görüntüle
			// - robotları düzenle
			// - robot ekle
			// - robotların durumlarını yazdır
			// - oyunu başlat
			// - oyundan çık
			options := options{
				displayHelpMessage,
				option{
					operation: func() GameState {
						selectRobotsThenConfigure(g.League.GetPlayers())
						return stateWaiting
					},
					description: "configure robots",
				},
				option{
					operation: func() GameState {
						addRobot(g.League.GetPlayers())
						return stateWaiting
					},
					description: "add a robot",
				},
				option{
					operation: func() GameState {
						generateReport(g.League.GetPlayers())
						return stateWaiting
					},
					description: "generate report",
				},
				option{
					operation: func() GameState {
						fmt.Println("starting the game")
						return stateMatchConfirmation
					},
					description: "start the game",
				},
				option{
					operation: func() GameState {
						fmt.Println("exiting the game")
						return stateGameOver
					},
					description: "exit game",
				},
			}
			g.SetState(askOptions(options)) // Seçenekleri kullanıcıya yönelt
		case stateMatchConfirmation: // Kullanıcıya ligi göster ve onayını al, onay vermezse yeni lig oluştur
			for _, v := range g.League.Matches { // Ligi yazdır
				fmt.Println(
					robot.GetRobotById(v.GetTeamA(), g.League.GetPlayers()).GetName(),
					"vs",
					robot.GetRobotById(v.GetTeamB(), g.League.GetPlayers()).GetName(),
				)
			}

			fmt.Println("would you like to continue?")
			returnedState := askOptions(options{
				option{
					operation: func() GameState {
						return statePlayingMatch // Oyun başladı durumuna geç
					},
					description: "yes",
				},
				option{
					operation: func() GameState {
						g.League.PopulateMatches() // Yeni lig oluştur
						return stateWaiting
					},
					description: "no",
				},
			})

			g.SetState(returnedState) // Oyunun durumunu değiştir
		case statePlayingMatch:
			// Kullanıcı onay verdiyse sırayla robotları karşı karşıya getir
			for _, val := range g.League.GetMatches() {
				if val.GetStatus() == robot_league.StatusNotPlayed {
					res := startMatch( // Robotları dövüştür
						robot.GetRobotById(val.GetTeamA(), g.League.GetPlayers()),
						robot.GetRobotById(val.GetTeamB(), g.League.GetPlayers()))

					val.SetStatus(robot_league.StatusPlayed)
					val.SetWinner(res) // Kazananı kaydet, daha sonra puan hesaplama vs. için kullanılabilir

					g.SetState(stateMatchEnded) // Maç bitti durumuna geç
					break
				}
			}
		case stateMatchEnded:
			gameEnded := g.League.Ended() // Bütün maçlar bitti mi kontrol et
			if gameEnded {
				g.SetState(stateGameOver) // Bittiyse oyunu bitir
			} else {
				g.SetState(statePlayingMatch) // Bitmediyse sonraki maça geç
			}
		case stateGameOver:
			// Maç bittiyse sonuçları yazdır ve çık
			if g.League.Ended() {
				for _, v := range g.League.Matches {
					fmt.Println(
						robot.GetRobotById(v.GetTeamA(), g.League.GetPlayers()).GetName(),
						"vs",
						robot.GetRobotById(v.GetTeamA(), g.League.GetPlayers()).GetName(),
						"winner: ",
						robot.GetRobotById(v.GetWinner(), g.League.GetPlayers()).GetName())
				}
			}
			os.Exit(0) // Programı sonlandır
		}
	}
}

// Yapıların değerlerini fonksiyona kopyalamak verimsiz olacağı için
// işaretçilerini veriyoruz
func startMatch(r1, r2 *robot.Robot) robot.RobotId {
	turnA := true // Sıra A'ncı yani birinci takımdan başlıyor
	var winner robot.RobotId
	for {
		if r1.GetHealth() < 0 || r2.GetHealth() < 0 {
			if r1.GetHealth() < 0 { // 1. oyuncunun canı kalmadıysa
				winner = r2.GetId()
				fmt.Println(r2.GetName() + " wins the round!\n\n")
				break
			} else if r2.GetHealth() < 0 { // 2. oyuncunun canı kalmadıysa
				winner = r1.GetId()
				fmt.Println(r1.GetName() + " wins the round!\n\n")
				break
			}
		} else {
			skillIndex := utils.Random3() // rastgele bir yetenek seç
			if turnA {
				r1.Attact(skillIndex, r2)
				// 1. robotu 2. ye saldırt ve sonuçlarını yazdır
				skillUsed := r1.GetSkill(skillIndex)
				fmt.Println(r1.GetName() + " deals " + strconv.FormatInt(-int64(skillUsed.GetHpEffect()), 10) + " damage")
				fmt.Println(r1.GetName() + "'s health: " + strconv.FormatInt(int64(r1.GetHealth()), 10))
				fmt.Println(r2.GetName() + "'s health: " + strconv.FormatInt(int64(r2.GetHealth()), 10))
				fmt.Print("------------\n")
				time.Sleep(time.Second * 3) // 3 saniye bekle
			} else {
				r2.Attact(skillIndex, r1)
				// 2. robotu 1. ye saldırt ve sonucu yazdır
				skillUsed := r2.GetSkill(skillIndex)
				fmt.Println(r2.GetName() + " deals " + strconv.FormatInt(-int64(skillUsed.GetHpEffect()), 10) + " damage")
				fmt.Println(r1.GetName() + "'s health: " + strconv.FormatInt(int64(r1.GetHealth()), 10))
				fmt.Println(r2.GetName() + "'s health: " + strconv.FormatInt(int64(r2.GetHealth()), 10))
				fmt.Print("------------\n")
				time.Sleep(time.Second * 3)
			}
			turnA = !turnA // Sıra diğer robota geçer
		}
	}

	// Maç bittikten sonra canları sıfırla
	r1.SetHealth(100)
	r2.SetHealth(100)
	return winner
}

// SETTERS
func (g *Game) SetState(s GameState) {
	g.state = s
}

// GETTERS
func (g *Game) GetState() GameState {
	return g.state
}
