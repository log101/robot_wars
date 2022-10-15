// Lig yapısı burada tanımlı
package robot_league

import (
	"errors"
	"math"

	"github.com/robot_wars/game/robot"
)

type Players map[robot.RobotId]*robot.Robot
type RobotLeague struct { // Robot ligi oyunculardan ve fikstürlerden oluşuyor
	players Players
	Matches map[MatchId]*MatchResult
}

type MatchStatus bool

// Oyun durumda yaptığımız gibi maçın da iki durumu olabilir
// oynanmış veya oynanmamış, bu bilgiyi ligin bitip bitmediğini
// tespit etmek için kullanacağız
const (
	StatusPlayed    MatchStatus = true
	StatusNotPlayed MatchStatus = false
)

type MatchId int

// Maç sonucu takımları, maçın oynanıp oynanmadığı bilgisini
// ve galibi tutuyor
type MatchResult struct {
	teamA  robot.RobotId
	teamB  robot.RobotId
	status MatchStatus
	winner robot.RobotId
}

// RobotLeague SETTERS
func (r *RobotLeague) SetPlayers(p Players) {
	r.players = p
}

func (r *RobotLeague) SetMatches(m map[MatchId]*MatchResult) {
	r.Matches = m
}

// League GETTERS
func (r *RobotLeague) GetPlayers() Players {
	return r.players
}

func (r *RobotLeague) GetMatches() map[MatchId]*MatchResult {
	return r.Matches
}

// MatchResult SETTERS
func (m *MatchResult) SetStatus(s MatchStatus) {
	m.status = s
}

func (m *MatchResult) SetWinner(id robot.RobotId) {
	m.winner = id
}

// MatchResult GETTERS
func (m *MatchResult) GetStatus() MatchStatus {
	return m.status
}

func (m *MatchResult) GetWinner() robot.RobotId {
	return m.winner
}

func (m *MatchResult) GetTeamA() robot.RobotId {
	return m.teamA
}

func (m *MatchResult) GetTeamB() robot.RobotId {
	return m.teamB
}

// Robotları map olarak tuttuğumuz için Lige kaydederken aynı eşleşmeyi bir daha (tersinden)
// yapmamak için bir eşleşme fonksiyonu kullanıyoruz, böylece daha önce bu eşleşmenin yapılıp
// yapılmadığını bulabiliriz
func pairingFunction(x, y robot.RobotId) MatchId {
	// Cantor Pairing Function
	// https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
	return MatchId((robot.RobotId(math.Pow(float64(x), 2)) + (3 * x) + (2 * x * y) + y + robot.RobotId(math.Pow(float64(y), 2))) / 2)
}

// Ligde oyuncuları birbiriyle eşleştiren metod
// Herkes birbiriyle bir sefer dövüşüyor
func (r *RobotLeague) PopulateMatches() error {
	r.Matches = make(map[MatchId]*MatchResult)

	if len(r.players) == 0 {
		return errors.New("teams are empty")
	} else if len(r.players) < 2 {
		return errors.New("number of teams should be more than 2")
	} else {
		for i1 := range r.players {
			for i2 := range r.players {
				matchId := pairingFunction(
					robot.RobotId(math.Min(float64(i1), float64(i2))),
					robot.RobotId(math.Max(float64(i1), float64(i2))))
				_, ok := r.Matches[matchId]
				if i1 == i2 || ok {
					continue
				}

				result := MatchResult{
					teamA:  robot.RobotId(math.Min(float64(i1), float64(i2))),
					teamB:  robot.RobotId(math.Max(float64(i1), float64(i2))),
					status: StatusNotPlayed,
				}

				r.Matches[matchId] = &result
			}
		}
	}

	return nil
}

// Ligin bitip bitmediğini kontrol etmek için
// maçların durumlarına bakıyoruz
func (r *RobotLeague) Ended() bool {
	for _, v := range r.Matches {
		if v.status == StatusNotPlayed {
			return false
		}
	}

	return true
}
