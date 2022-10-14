package robot

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestRobotLeague_PopulateMatches(t *testing.T) {

	got := maps.Equal[map[RobotId]Robot](StarterRobots, StarterRobots)
	if got != true {
		t.Errorf("maps must be equal")
	}
}
