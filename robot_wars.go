package main

import (
	"github.com/robot_wars/game"
	"github.com/robot_wars/game/robot"
	"github.com/robot_wars/game/robot_league"
)

// Uygulamanın başlangıç noktası:
// game.go'da tanımlanmış fonksiyona robotların olduğu bir lig
// tanımladıktan sonra eşleşmeleri oluşturuyoruz
func main() {
	var myLeague robot_league.RobotLeague
	myLeague.SetPlayers(robot.StarterRobotsSample) // seed dosyasından robotları yükle
	myLeague.PopulateMatches()                     // Eşleşmeleri oluştur

	myGame := game.Game{
		League: &myLeague,
	}

	myGame.Start()
}
