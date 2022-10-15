package robot

import (
	"testing"
)

func TestGetRobotById(t *testing.T) {
	t.Run("can get robots with id", func(t *testing.T) {
		got := GetRobotById(RobotId(1), StarterRobots)

		expected := StarterRobots[RobotId(1)]

		t.Log(&got)
		t.Log(&expected)

		if got != expected {
			t.Error("cannot get robots by id")
		}
	})
}

func TestAttact(t *testing.T) {
	t.Run("robot can attack other robots", func(t *testing.T) {
		robot1 := StarterRobots[RobotId(1)]
		robot2 := StarterRobots[RobotId(2)]

		deafultHealth := robot2.health
		robot1.Attact(1, robot2)

		got := robot2.health
		expected := deafultHealth + robot1.skills[1].hpEffect

		if got != expected {
			t.Log(expected)
			t.Log(got)
			t.Error("robot cannot attack")
		}
	})
}
