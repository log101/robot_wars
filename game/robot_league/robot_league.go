package robot_league

import (
	"errors"
	"math"

	"github.com/robot_wars/game/robot"
)

type RobotLeague struct {
	players Players
	matches map[MatchId]MatchResult
}

type MatchStatus bool

const (
	statusPlayed    MatchStatus = true
	statusNotPlayed MatchStatus = false
)

type MatchId int
type Players map[robot.RobotId]robot.Robot
type MatchResult struct {
	teamA  robot.RobotId
	teamB  robot.RobotId
	done   MatchStatus
	winner robot.RobotId
}

// RobotLeague SETTERS
func (r *RobotLeague) SetPlayers(p Players) {
	r.players = p
}

func (r *RobotLeague) SetMatches(m map[MatchId]MatchResult) {
	r.matches = m
}

// League GETTERS
func (r *RobotLeague) GetPlayers() Players {
	return r.players
}

func (r *RobotLeague) GetMatches() map[MatchId]MatchResult {
	return r.matches
}

// MatchResult SETTERS
func (m *MatchResult) SetDone() {
	m.done = statusPlayed
}

func (m *MatchResult) SetWinner(id robot.RobotId) {
	m.winner = id
}

// MatchResult GETTERS
func (m *MatchResult) GetDone() MatchStatus {
	return m.done
}

func (m *MatchResult) GetWinner() robot.RobotId {
	return m.winner
}

// Robotları map olarak tuttuğumuz için Lige kaydederken aynı eşleşmeyi bir daha (tersinden)
// yapmamak için bir eşleşme fonksiyonu kullanıyoruz, böylece daha önce bu eşleşmenin yapılıp
// yapılmadığını bulabiliriz
func pairingFunction(x, y robot.RobotId) MatchId {
	// Cantor Pairing Function
	// https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
	return MatchId((robot.RobotId(math.Pow(float64(x), 2)) + (3 * x) + (2 * x * y) + y + robot.RobotId(math.Pow(float64(y), 2))) / 2)
}

func (r *RobotLeague) PopulateMatches() error {
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
				_, ok := r.matches[matchId]
				if i1 == i2 || ok {
					continue
				}

				r.matches[matchId] = MatchResult{
					teamA: robot.RobotId(math.Min(float64(i1), float64(i2))),
					teamB: robot.RobotId(math.Max(float64(i1), float64(i2))),
					done:  statusNotPlayed,
				}
			}
		}
	}

	return nil
}

func (r *RobotLeague) ended() bool {
	for _, v := range r.matches {
		if v.done == statusNotPlayed {
			return false
		}
	}

	return true
}
