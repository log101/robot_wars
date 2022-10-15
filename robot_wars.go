package main

import (
	"github.com/robot_wars/game"
	"github.com/robot_wars/game/robot"
	"github.com/robot_wars/game/robot_league"
)

func main() {
	var myLeague robot_league.RobotLeague
	myLeague.SetPlayers(robot.StarterRobotsSample)
	myLeague.PopulateMatches()

	myGame := game.Game{
		League: &myLeague,
	}

	myGame.Start()
}
