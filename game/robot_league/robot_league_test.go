package robot_league

import (
	"testing"

	"golang.org/x/exp/maps"

	"github.com/robot_wars/game/robot"
)

func TestRobotLeague_League(t *testing.T) {
	sampleMatches := make(map[MatchId]*MatchResult)

	sampleLeague := RobotLeague{
		players: robot.StarterRobotsSample,
		Matches: sampleMatches,
	}

	sampleLeague.PopulateMatches()

	t.Run("number of matches are correct", func(t *testing.T) {
		playerCount := len(sampleLeague.players)
		if len(sampleLeague.Matches) != (playerCount*(playerCount-1))/2 {
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

		expectedMatches := map[MatchId]MatchResult{
			correctPairs[0]: {
				teamA:  1,
				teamB:  2,
				status: StatusNotPlayed,
			},
			correctPairs[1]: {
				teamA:  1,
				teamB:  3,
				status: StatusNotPlayed,
			},
			correctPairs[2]: {
				teamA:  1,
				teamB:  4,
				status: StatusNotPlayed,
			},
			correctPairs[3]: {
				teamA:  2,
				teamB:  3,
				status: StatusNotPlayed,
			},
			correctPairs[4]: {
				teamA:  2,
				teamB:  4,
				status: StatusNotPlayed,
			},
			correctPairs[5]: {
				teamA:  3,
				teamB:  4,
				status: StatusNotPlayed,
			},
		}

		leagueMap := make(map[MatchId]MatchResult)

		for k, v := range sampleLeague.Matches {
			leagueMap[k] = *v
		}

		if !maps.Equal(expectedMatches, leagueMap) {
			t.Log(expectedMatches)
			t.Log(sampleLeague.Matches)
			t.Error("everyone is not matched")
		}
	})

	t.Run("ended returns false when league is not finished", func(t *testing.T) {
		got := sampleLeague.Ended()

		expected := false

		if got != expected {
			t.Error("should return false, league is not ended")
		}
	})

	t.Run("ended returns true when league is finished", func(t *testing.T) {
		samplePlayers := Players{
			1: sampleLeague.players[1],
			2: sampleLeague.players[2],
			3: sampleLeague.players[3],
		}

		sampleMatches := map[MatchId]*MatchResult{
			pairingFunction(1, 2): {
				teamA:  1,
				teamB:  2,
				status: StatusPlayed,
				winner: 1,
			},
			pairingFunction(1, 3): {
				teamA:  1,
				teamB:  3,
				status: StatusPlayed,
				winner: 1,
			},
			pairingFunction(2, 3): {
				teamA:  2,
				teamB:  3,
				status: StatusPlayed,
				winner: 2,
			},
		}

		finishedLeague := RobotLeague{
			players: samplePlayers,
			Matches: sampleMatches,
		}

		got := finishedLeague.Ended()

		expected := true

		if got != expected {
			t.Error("should return true, league has ended")
		}
	})

	t.Run("league players' getters and setters are working", func(t *testing.T) {
		sampleLeague.SetPlayers(robot.StarterRobots)

		gotPlayers := sampleLeague.GetPlayers()
		expectedPlayers := robot.StarterRobots

		if !maps.Equal(gotPlayers, expectedPlayers) {
			t.Error("error in setting or getting players")
		}
	})

	t.Run("league matches' getters and setters are working", func(t *testing.T) {
		expectedMatches := map[MatchId]*MatchResult{
			1: {
				teamA:  1,
				teamB:  2,
				status: StatusNotPlayed,
				winner: 1,
			},
			2: {
				teamA:  3,
				teamB:  4,
				status: StatusPlayed,
				winner: 3,
			},
		}

		sampleLeague.SetMatches(expectedMatches)

		gotMatches := sampleLeague.GetMatches()

		if !maps.Equal(gotMatches, expectedMatches) {
			t.Error("error in setting or getting matches")
		}

	})
}

func TestRobotLeague_MatchResult(t *testing.T) {
	sampleMatchResult := MatchResult{
		teamA:  1,
		teamB:  2,
		status: StatusNotPlayed,
	}

	t.Run("setters and getters are working", func(t *testing.T) {
		expectedDone := StatusPlayed
		sampleMatchResult.SetStatus(StatusPlayed)
		gotDone := sampleMatchResult.GetStatus()

		if expectedDone != gotDone {
			t.Error("error while setting or getting match result done")
		}

		expectedWinner := robot.RobotId(1)
		sampleMatchResult.SetStatus(StatusPlayed)
		sampleMatchResult.SetWinner(expectedWinner)

		gotWinner := sampleMatchResult.GetWinner()

		if expectedWinner != gotWinner {
			t.Error("error while setting or getting match result")
		}

	})

}
