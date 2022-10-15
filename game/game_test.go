package game

import (
	"testing"

	"github.com/robot_wars/game/robot"
)

func Test_startMatch(t *testing.T) {
	t.Run("match ends with a winner robot", func(t *testing.T) {
		r1 := robot.StarterRobots[1]
		r2 := robot.StarterRobots[2]

		got := startMatch(r1, r2)

		if got > 2 {
			t.Error("cant finish the match")
		}
	})
}
