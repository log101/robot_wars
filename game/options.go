// Kullanıcıya sorulacak soruların tanımlandığı yer
// Diyaloglarda diyebiliriz
package game

import (
	"fmt"

	"github.com/robot_wars/game/robot"
	league "github.com/robot_wars/game/robot_league"
)

// Soru/seçenek yapısı
// Her seçenek bir açıklama ve fonksiyondan oluşuyor
// Seçenek seçildiği takdirde fonksiyon çalışıyor
// Geri döndürülen "Oyun Durumu" değişkeni ile oyunumuzun
// durumunda(state) değişiklik yapabiliyoruz
// İnteraktif, kullanıcıdan dönüt alınması gereken
// neredeyse tüm özellikler bu fonksiyon yardımıyla oluşturuldu
type option struct {
	operation   func() GameState // Geri oyun durumu döndüren fonksiyon
	description string           // Seçeneğin açıklaması
}

// Kodun okunabilirliğini arttırmak adına
// bu tarz tip tanımlamalarından çokça yaptım
type options []option

// Kullanıcıya seçenekler sun
func askOptions(ops options) GameState {
	for i, opt := range ops { // Seçeneklerin açıklamarını ve indekslerini yazdır
		fmt.Println(i+1, ")", opt.description)
	}

	var answer int
	for { // Geçerli yanıt alana kadar sor
		fmt.Print("Please Select an option\n> ")
		fmt.Scanf("%d", &answer)
		if (answer > len(ops)) || (answer < 0) {
			fmt.Println("invalid option") // Geçersiz yanıtları reddet
		} else {
			break
		}
	}

	return ops[answer-1].operation() // Seçilen seçeneğin fonksiyonunu çalıştır
}

// Proje niteliklerinde belirtilen yardım mesajını görüntüle
// seçenek olduğu için tekrar tekrar kullanılabilir
var displayHelpMessage option = option{
	operation: func() GameState {
		fmt.Println(`
		This is a randomly generated robot figthing game, please select the option you prefer!
		You can add new robots or modify the existing ones!
		Have fun!
		`)
		return stateWaiting
	},
	description: "display help message",
}

// Ligde bulunan robotların isimlerini, canlarını ve yeteneklerini yazdır
func generateReport(p league.Players) {
	for _, v := range p {
		fmt.Println(
			"Name:", v.GetName(),
			"Health:", v.GetHealth(),
			"Skill 1:", v.GetSkillFeatures(0),
			"Skill 2:", v.GetSkillFeatures(1),
			"Skill 3:", v.GetSkillFeatures(2),
		)
	}
}

// Robotlar üzerinde değişlik yap
func selectRobotsThenConfigure(p league.Players) {
	for i, v := range p {
		fmt.Println(i, ")", v.GetName())
	}

	fmt.Println("Please select a robot!")

	var selection int
	fmt.Scanf("%d", &selection)

	selectedRobot, ok := p[robot.RobotId(selection)]
	if ok {
		configureRobot(selectedRobot)
	}
}

// İsim, can ve yetenekler üzerinde değişiklik yap
func configureRobot(r *robot.Robot) {
	askOptions(options{
		option{
			operation: func() GameState {
				fmt.Println("enter a new name")

				var newName string
				fmt.Scanf("%s", &newName)

				r.SetName(newName)
				fmt.Println("name successfully changed")

				return stateWaiting
			},
			description: "change name",
		},
		option{
			operation: func() GameState {
				fmt.Println("enter a new health value")

				var newHealth int
				fmt.Scanf("%d", &newHealth)

				r.SetHealth(newHealth)
				fmt.Println("health successfully changed")

				return stateWaiting
			},
			description: "change health",
		},
		option{
			operation: func() GameState {
				for i, k := range robot.StarterSkills {
					fmt.Println(i+1, ")", k.GetName(), " power:", -k.GetHpEffect())
				}
				fmt.Println("###################")
				fmt.Println("select a skill from the list")

				var selection int
				fmt.Scanf("%d", &selection)

				r.SetSkill(0, robot.StarterSkills[selection-1])
				fmt.Println("skill successfully changed!")

				return stateWaiting
			},
			description: "change first skill",
		},
		option{
			operation: func() GameState {
				for i, k := range robot.StarterSkills {
					fmt.Println(i+1, ")", k.GetName(), " power:", -k.GetHpEffect())
				}
				fmt.Println("###################")
				fmt.Println("select a skill from the list")

				var selection int
				fmt.Scanf("%d", &selection)

				r.SetSkill(1, robot.StarterSkills[selection-1])
				fmt.Println("skill successfully changed!")

				return stateWaiting
			},
			description: "change second skill",
		},
		option{
			operation: func() GameState {
				for i, k := range robot.StarterSkills {
					fmt.Println(i+1, ")", k.GetName(), " power:", -k.GetHpEffect())
				}
				fmt.Println("###################")
				fmt.Println("select a skill from the list")

				var selection int
				fmt.Scanf("%d", &selection)

				r.SetSkill(2, robot.StarterSkills[selection-1])
				fmt.Println("skill successfully changed!")

				return stateWaiting
			},
			description: "change third skill",
		},
	})
}

// Lige yeni bir robot ekle
func addRobot(p league.Players) {
	var (
		name   string
		skills [3]robot.Skill
	)

	fmt.Println("enter a name for the robot!")
	fmt.Scanf("%s", &name)

	for i := 0; i < 3; i++ { // 3 tane yetenek seç
		var selection int
		fmt.Println("choose a skill from the set")
		for k, v := range robot.StarterSkills {
			fmt.Println(k+1, ")", v.GetName(), "power:", v.GetHpEffect())
		}
		fmt.Scanf("%d", &selection)
		skills[i] = robot.StarterSkills[selection-1]
		fmt.Println("skill successfully added")
	}

	newRobot := robot.CreateRobot(robot.RobotId(len(p)+1), name, 100, skills) // Seçilen robotu oluştur
	p[newRobot.GetId()] = newRobot

	// Robot oluşturulduktan sonra olumlu geribildirimde bulun
	fmt.Println("robot", newRobot.GetName(), "successfully created")
}
