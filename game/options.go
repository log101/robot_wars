package game

import (
	"fmt"

	"github.com/robot_wars/game/robot"
	league "github.com/robot_wars/game/robot_league"
)

type option struct {
	operation   func() GameState
	description string
}

type options []option

func askOptions(ops options) GameState {
	for i, opt := range ops {
		fmt.Println(i+1, ")", opt.description)
	}

	var answer int
	for {
		fmt.Print("Please Select an option\n> ")
		fmt.Scanf("%d", &answer)
		if (answer > len(ops)) || (answer < 0) {
			fmt.Println("invalid option")
		} else {
			break
		}
	}

	return ops[answer-1].operation()
}

var displayHelpMessage option = option{
	operation: func() GameState {
		fmt.Println("go run .")
		return stateWaiting
	},
	description: "display help message",
}

func generateReport(p league.Players) {
	for _, v := range p {
		fmt.Println(
			"Name:", v.GetName(),
			"Health:", v.GetHealth(),
			"Skill 1:", v.GetSkillFeatures(0),
			"Skill 2:", v.GetSkillFeatures(1),
			"Skill 3:", v.GetSkillFeatures(2),
		)
	}
}

func selectRobotsThenConfigure(p league.Players) {
	for i, v := range p {
		fmt.Println(i, ")", v.GetName())
	}

	fmt.Println("Please select a robot!")

	var selection int
	fmt.Scanf("%d", &selection)

	selectedRobot, ok := p[robot.RobotId(selection)]
	if ok {
		configureRobot(selectedRobot)
	}
}

func configureRobot(r *robot.Robot) {
	askOptions(options{
		option{
			operation: func() GameState {
				fmt.Println("enter a new name")

				var newName string
				fmt.Scanf("%s", newName)

				r.SetName(newName)
				fmt.Println("name successfully changed")

				return stateWaiting
			},
			description: "change name",
		},
		option{
			operation: func() GameState {
				fmt.Println("enter a new health value")

				var newHealth int
				fmt.Scanf("%d", &newHealth)

				r.SetHealth(newHealth)
				fmt.Println("health successfully changed")

				return stateWaiting
			},
			description: "change health",
		},
		option{
			operation: func() GameState {
				return stateWaiting
			},
			description: "change first skill",
		},
		option{
			operation: func() GameState {
				return stateWaiting
			},
			description: "change second skill",
		},
		option{
			operation: func() GameState {
				return stateWaiting
			},
			description: "change third skill",
		},
	})
}
