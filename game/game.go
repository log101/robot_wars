package game

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/robot_wars/game/robot"
	"github.com/robot_wars/game/robot_league"
	"github.com/robot_wars/utils"
)

type GameState string

var (
	stateWaiting           GameState = "stateWaiting"
	statePlayingMatch      GameState = "statePlayingMatch"
	stateMatchConfirmation GameState = "stateMatchConfirmation"
	stateMatchEnded        GameState = "stateMatchEnded"
	stateGameOver          GameState = "stateGameOver"
)

type Game struct {
	League *robot_league.RobotLeague
	state  GameState
}

// METHODS
func (g *Game) Start() {
	g.SetState(stateWaiting)
	displayHelpMessage.operation()

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
						addRobot(g.League.GetPlayers())
						return stateWaiting
					},
					description: "add a robot",
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
						return stateMatchConfirmation
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
		case stateMatchConfirmation:
			for _, v := range g.League.Matches {
				fmt.Println(
					robot.GetRobotById(v.GetTeamA(), g.League.GetPlayers()).GetName(),
					"vs",
					robot.GetRobotById(v.GetTeamB(), g.League.GetPlayers()).GetName(),
				)
			}

			fmt.Println("would you like to continue?")
			returnedState := askOptions(options{
				option{
					operation: func() GameState {
						return statePlayingMatch
					},
					description: "yes",
				},
				option{
					operation: func() GameState {
						g.League.PopulateMatches()
						return stateWaiting
					},
					description: "no",
				},
			})

			g.SetState(returnedState)
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
				fmt.Println(
					robot.GetRobotById(v.GetTeamA(), g.League.GetPlayers()).GetName(),
					"vs",
					robot.GetRobotById(v.GetTeamA(), g.League.GetPlayers()).GetName(),
					"winner: ",
					robot.GetRobotById(v.GetWinner(), g.League.GetPlayers()).GetName())
			}
			os.Exit(0)
		}
	}
}

// Koca koca yapıları fonksiyona kopyalamak verimsiz olacağı için
// işaretçilerini veriyoruz
func startMatch(r1, r2 *robot.Robot) robot.RobotId {
	turnA := true // Sıra A'ncı yani birinci takımdan başlıyor
	var winner robot.RobotId
	for {
		if r1.GetHealth() < 0 || r2.GetHealth() < 0 {
			if r1.GetHealth() < 0 {
				winner = r2.GetId()
				fmt.Println(r2.GetName() + " wins the round!\n\n")
				break
			} else if r2.GetHealth() < 0 {
				winner = r1.GetId()
				fmt.Println(r1.GetName() + " wins the round!\n\n")
				break
			}
		} else {
			skillIndex := utils.Random3()
			if turnA {
				r1.Attact(skillIndex, r2)
				skillUsed := r1.GetSkill(skillIndex)
				fmt.Println(r1.GetName() + " deals " + strconv.FormatInt(-int64(skillUsed.GetHpEffect()), 10) + " damage")
				fmt.Println(r1.GetName() + "'s health: " + strconv.FormatInt(int64(r1.GetHealth()), 10))
				fmt.Println(r2.GetName() + "'s health: " + strconv.FormatInt(int64(r2.GetHealth()), 10))
				fmt.Print("------------\n")
				time.Sleep(time.Second * 3)
			} else {
				r2.Attact(skillIndex, r1)
				skillUsed := r2.GetSkill(skillIndex)
				fmt.Println(r2.GetName() + " deals " + strconv.FormatInt(-int64(skillUsed.GetHpEffect()), 10) + " damage")
				fmt.Println(r1.GetName() + "'s health: " + strconv.FormatInt(int64(r1.GetHealth()), 10))
				fmt.Println(r2.GetName() + "'s health: " + strconv.FormatInt(int64(r2.GetHealth()), 10))
				fmt.Print("------------\n")
				time.Sleep(time.Second * 3)
			}
			turnA = !turnA
		}
	}

	r1.SetHealth(100)
	r2.SetHealth(100)
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
