package league

import (
	"errors"

	"github.com/robot_wars/robot"
)

type MatchStatus bool

const (
	statusPlayed    MatchStatus = true
	statusNotPlayed MatchStatus = false
)

type Players []robot.Robot

type Match struct {
	teamA  robot.RobotId
	teamB  robot.RobotId
	done   MatchStatus
	winner robot.RobotId
}

type League struct {
	players Players
	matches []Match
}

func (l *League) populateMatches() error {
	if len(l.players) == 0 {
		return errors.New("teams are empty")
	} else if len(l.players) < 2 {
		return errors.New("number of teams should be more than 2")
	} else {
		for i1, p1 := range l.players {
			for i2, p2 := range l.players {
				if i1 == i2 {
					continue
				}

				newMatch := Match{
					teamA: p1.GetId(),
					teamB: p2.GetId(),
					done:  statusNotPlayed,
				}

				l.matches = append(l.matches, newMatch)
			}
		}
	}

	return nil
}
