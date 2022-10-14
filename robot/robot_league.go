package robot

import (
	"errors"
	"math"
)

type MatchStatus bool
type MatchId int

const (
	statusPlayed    MatchStatus = true
	statusNotPlayed MatchStatus = false
)

type Players map[RobotId]Robot

type MatchResult struct {
	teamA  RobotId
	teamB  RobotId
	done   MatchStatus
	winner RobotId
}

type RobotLeague struct {
	players Players
	matches map[MatchId]MatchResult
}

// Robotları map olarak tuttuğumuz için Lige kaydederken aynı eşleşmeyi bir daha (tersinden)
// yapmamak için bir eşleşme fonksiyonu kullanıyoruz, böylece daha önce bu eşleşmenin yapılıp
// yapılmadığını bulabiliriz
func pairingFunction(x, y RobotId) MatchId {
	// Cantor Pairing Function
	// https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
	return MatchId((RobotId(math.Pow(float64(x), 2)) + (3 * x) + (2 * x * y) + y + RobotId(math.Pow(float64(y), 2))) / 2)
}

func (r *RobotLeague) PopulateMatches() error {
	if len(r.players) == 0 {
		return errors.New("teams are empty")
	} else if len(r.players) < 2 {
		return errors.New("number of teams should be more than 2")
	} else {
		for i1, _ := range r.players {
			for i2, _ := range r.players {
				matchId := pairingFunction(
					RobotId(math.Min(float64(i1), float64(i2))),
					RobotId(math.Max(float64(i1), float64(i2))))
				_, ok := r.matches[matchId]
				if i1 == i2 || ok {
					continue
				}

				r.matches[matchId] = MatchResult{
					teamA: RobotId(math.Min(float64(i1), float64(i2))),
					teamB: RobotId(math.Max(float64(i1), float64(i2))),
					done:  statusNotPlayed,
				}
			}
		}
	}

	return nil
}
