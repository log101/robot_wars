package game

import (
	"fmt"
	"os"

	"github.com/robot_wars/game/robot"
	"github.com/robot_wars/game/robot_league"
	"github.com/robot_wars/utils"
)

type GameState string

var (
	stateWaiting       GameState = "stateWaiting"
	stateCreatingRobot GameState = "stateCreatingRobot"
	statePlayingMatch  GameState = "statePlayingMatch"
	stateMatchEnded    GameState = "stateMatchEnded"
	stateGameOver      GameState = "stateGameOver"
)

type Game struct {
	League *robot_league.RobotLeague
	state  GameState
}

// METHODS
func (g *Game) Start() {
	g.SetState(stateWaiting)

	for {
		switch g.GetState() {
		case stateWaiting:
			options := options{
				displayHelpMessage,
				option{
					operation: func() GameState {
						selectRobotsThenConfigure(g.League.GetPlayers())
						return stateWaiting
					},
					description: "configure robots",
				},
				option{
					operation: func() GameState {
						generateReport(g.League.GetPlayers())
						return stateWaiting
					},
					description: "generate report",
				},
				option{
					operation: func() GameState {
						fmt.Println("starting the game")
						return statePlayingMatch
					},
					description: "start the game",
				},
				option{
					operation: func() GameState {
						fmt.Println("exiting the game")
						return stateGameOver
					},
					description: "exit game",
				},
			}
			g.SetState(askOptions(options))
		case stateCreatingRobot:
		case statePlayingMatch:
			for _, val := range g.League.GetMatches() {
				if val.GetStatus() == robot_league.StatusNotPlayed {
					res := startMatch(
						robot.GetRobotById(val.GetTeamA(), g.League.GetPlayers()),
						robot.GetRobotById(val.GetTeamB(), g.League.GetPlayers()))

					val.SetStatus(robot_league.StatusPlayed)
					val.SetWinner(res)

					g.SetState(stateMatchEnded)
					break
				}
			}
		case stateMatchEnded:
			gameEnded := g.League.Ended()
			if gameEnded {
				g.SetState(stateGameOver)
			} else {
				g.SetState(statePlayingMatch)
			}
		case stateGameOver:
			for _, v := range g.League.Matches {
				fmt.Println(v.GetTeamA(), "vs", v.GetTeamB(), "winner: ", v.GetWinner())
			}
			os.Exit(0)
		}
	}
}

// Koca koca yapıları fonksiyona kopyalamak verimsiz olacağı için
// işaretçilerini veriyoruz
func startMatch(r1, r2 *robot.Robot) robot.RobotId {
	turnA := true // Sıra A'ncı yani birinci takımdan başlıyor
	winner := robot.RobotId(0)
	for {
		if r1.GetHealth() < 0 {
			winner = r2.GetId()
			break
		} else if r2.GetHealth() < 0 {
			winner = r1.GetId()
			break
		} else {
			skillIndex := utils.Random3()
			if turnA {
				r1.Attact(skillIndex, r2)
			} else {
				r2.Attact(skillIndex, r1)
			}
		}
	}

	return winner
}

// SETTERS
func (g *Game) SetState(s GameState) {
	g.state = s
}

// GETTERS
func (g *Game) GetState() GameState {
	return g.state
}
