package robot

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestRobotLeague_PopulateMatches(t *testing.T) {
	sampleMatches := make(map[MatchId]MatchResult)

	sampleLeague := RobotLeague{
		players: StarterRobotsSample,
		matches: sampleMatches,
	}

	sampleLeague.PopulateMatches()

	t.Run("number of matches are correct", func(t *testing.T) {
		playerCount := len(sampleLeague.players)
		if len(sampleLeague.matches) != (playerCount*(playerCount-1))/2 {
			t.Error("number of matches are incorrectly calculated")
		}
	})

	// Herkes herkesle bir sefer karşılaşmış olmalı
	t.Run("everyone is matched", func(t *testing.T) {
		correctPairs := []MatchId{
			pairingFunction(1, 2),
			pairingFunction(1, 3),
			pairingFunction(1, 4),
			pairingFunction(2, 3),
			pairingFunction(2, 4),
			pairingFunction(3, 4),
		}

		expectedMatches := make(map[MatchId]MatchResult)
		expectedMatches = map[MatchId]MatchResult{
			correctPairs[0]: {
				teamA: 1,
				teamB: 2,
				done:  statusNotPlayed,
			},
			correctPairs[1]: {
				teamA: 1,
				teamB: 3,
				done:  statusNotPlayed,
			},
			correctPairs[2]: {
				teamA: 1,
				teamB: 4,
				done:  statusNotPlayed,
			},
			correctPairs[3]: {
				teamA: 2,
				teamB: 3,
				done:  statusNotPlayed,
			},
			correctPairs[4]: {
				teamA: 2,
				teamB: 4,
				done:  statusNotPlayed,
			},
			correctPairs[5]: {
				teamA: 3,
				teamB: 4,
				done:  statusNotPlayed,
			},
		}

		if !maps.Equal(expectedMatches, sampleLeague.matches) {
			t.Log(expectedMatches)
			t.Log(sampleLeague.matches)
			t.Error("everyone is not matched")
		}
	})
}
