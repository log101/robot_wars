package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//------
//robotun genel özellikleri
type Robot struct {
	Name            string
	heal            int16
	speed           uint8
	physicalPower   uint8
	width           uint8
	length          uint8
	height          uint8
	volume          uint16
	dodgePossibilty uint8
	isDodgeSkill    bool
}

//atak kabiliyeti
func (r Robot) AtackSkill(enemyRobot Robot) (Robot, Robot) {
	randomNumber := uint8(rand.Intn(7)) + enemyRobot.dodgePossibilty/3
	if randomNumber < 7 {
		fmt.Println(r.Name, r.physicalPower, "hasar vurdu")
		enemyRobot.heal = enemyRobot.heal - int16(r.physicalPower)
	} else {
		fmt.Println(enemyRobot.Name, " sıyrıldı")
	}
	if r.isDodgeSkill == true {
		r.dodgePossibilty /= 2
	}
	r.isDodgeSkill = false
	return r, enemyRobot
}

//can basma kabiliyeti
func (r Robot) HealSkill() Robot {
	if r.isDodgeSkill == true {
		r.dodgePossibilty /= 2
	}
	r.heal += 25
	r.isDodgeSkill = false
	fmt.Println(r.Name, "Can Bastı")
	return r
}

//kaçınma kabiliyeti
func (r Robot) DodgeSkill() Robot {
	if r.isDodgeSkill == true {
		r.dodgePossibilty /= 2
	}
	r.dodgePossibilty *= 2
	r.isDodgeSkill = true
	fmt.Println(r.Name, "kaçış ihtimalini yükseltti")
	return r
}

func (r Robot) SetRobotProperty(name string) Robot {
	r.Name = name

	r.heal = int16(rand.Intn(21) + 80)          // heal
	r.speed = uint8(rand.Intn(5) + 1)           // speed
	r.physicalPower = uint8(rand.Intn(21) + 20) // physical power
	r.width = uint8(rand.Intn(15) + 15)         // width
	r.length = uint8(rand.Intn(15) + 15)        // length
	r.height = uint8(rand.Intn(15) + 15)        // height
	r.volume = uint16(r.width) * uint16(r.length) * uint16(r.height)
	r.dodgePossibilty = uint8(10000/r.volume + uint16(r.speed))
	return r
}

//------

func main() {
	var chosen int8
	var robots []Robot
	robots = createRobot(robots, 10, "")
	robots = createYourRobot(robots)
OuterLoop:
	for {
		fmt.Print("\nMENÜ \n------- \n0-Menü \n1-Kendi robotun \n2-Tüm robotlar \n3-Ligler \n4-Sıradaki rakip \n5-Dövüş \n6-Oyundan Çık\n")
	UserInputLoop:
		for {
			fmt.Scanf("%d\n", &chosen)
			switch chosen {
			case 0:
				{
					break UserInputLoop
				}
			case 1:
				{
					printRobotList(robots, true)
					chosen = -1
				}
			case 2:
				{
					printRobotList(robots, false)
					chosen = -1
				}
			case 4:
				{
					fmt.Println("Sıradaki rakip", robots[0])
				}
			case 5:
				{
					Fight(robots[len(robots)-1], robots[0], true)
				}
			case 6:
				{
					fmt.Println("Görüşmek üzere...")
					break OuterLoop
				}
			default:
				{
					fmt.Println("Hatalı sayı girişi tekrar girin")
				}

			}
		}
	}
}

//--------
//Robot oluşturma ve ekrana bilgilerini yazdırma
func createRobot(robots []Robot, countOfRobot int, name string) []Robot {

	for i := 0; i < countOfRobot; i++ {
		if name == "" {
			name = "robot" + strconv.Itoa(i+1)
		}
		var newRobot Robot
		newRobot = newRobot.SetRobotProperty(name)
		robots = append(robots, newRobot)
		name = ""
	}
	return robots
}

func printRobotList(robots []Robot, isPlayerRobot bool) {
	if isPlayerRobot {
		fmt.Println("İsmi: ", robots[len(robots)-1].Name)
		fmt.Println("Canı: ", robots[len(robots)-1].heal)
		fmt.Println("Fiziksel Gücü: ", robots[len(robots)-1].physicalPower)
		fmt.Println("Kaçınma Olasılığı: ", robots[len(robots)-1].dodgePossibilty)
		fmt.Println("---------")
	} else {
		for i := 0; i < len(robots); i++ {
			fmt.Println("İsmi: ", robots[i].Name)
			fmt.Println("Canı: ", robots[i].heal)
			fmt.Println("Fiziksel Gücü: ", robots[i].physicalPower)
			fmt.Println("Kaçınma Olasılığı: ", robots[i].dodgePossibilty)
			fmt.Println("---------")
		}
	}
}

func createYourRobot(robots []Robot) []Robot {
	var playerRobot Robot
	fmt.Print("Robotunun ismi : ")
	fmt.Scanf("%s\n", &playerRobot.Name)
	robots = createRobot(robots, 1, playerRobot.Name)
	return robots
}

//--------

//--------
//Arena kısmı

func Fight(robot1 Robot, robot2 Robot, isPlayerRobot bool) {
	var playerChosenSkill uint8
	fmt.Println("Sol Köşede: ", robot1.Name)
	fmt.Println("Sağ Köşede: ", robot2.Name)
	for {
		if !isPlayerRobot {
			time.Sleep(5 * time.Second)
			robot1, robot2 = robotUseSkill(robot1, robot2, 0)

			fmt.Println("İsmi: ", robot1.Name)
			fmt.Println("Canı: ", robot1.heal)
			fmt.Println("Kaçınma Olasılığı: ", robot1.dodgePossibilty)
			fmt.Println("---------")
		} else {
			fmt.Println("Hangi Skill'i kullanmak istiyorsun 1-Attack 2-Heal 3-DoubleDodge")
			for {
				fmt.Scanf("%d\n", &playerChosenSkill)
				if playerChosenSkill < 0 || playerChosenSkill > 3 {
					fmt.Println("Yanlış sayı girdiniz lütfen 1,2 veya 3 den bir tanesini giriniz.")
				} else {
					break
				}
			}
			robot1, robot2 = robotUseSkill(robot1, robot2, playerChosenSkill)

			fmt.Println("İsmi: ", robot1.Name)
			fmt.Println("Canı: ", robot1.heal)
			fmt.Println("Kaçınma Olasılığı: ", robot1.dodgePossibilty)
			fmt.Println("---------")
		}

		if isExplosion(robot2) {
			fmt.Println(robot2.Name, " Patladı")
			break
		}
		fmt.Println("Rakip bekleniyor")
		time.Sleep(5 * time.Second)
		robot2, robot1 = robotUseSkill(robot2, robot1, 0)
		if isExplosion(robot1) {
			fmt.Println(robot1.Name, " Patladı")
			break
		}
		fmt.Println("İsmi: ", robot2.Name)
		fmt.Println("Canı: ", robot2.heal)
		fmt.Println("Kaçınma Olasılığı: ", robot2.dodgePossibilty)
		fmt.Println("---------")
	}
}

func robotUseSkill(robot1 Robot, robot2 Robot, playerChosenSkill uint8) (Robot, Robot) {

	if playerChosenSkill == 0 {
		diceValue := rollToDice(25)

		if diceValue <= 15 {
			return robot1.AtackSkill(robot2)
		} else if diceValue <= 20 {
			return robot1.HealSkill(), robot2
		} else {
			return robot1.DodgeSkill(), robot2
		}
	} else {
		if playerChosenSkill == 1 {
			return robot1.AtackSkill(robot2)
		} else if playerChosenSkill == 2 {
			return robot1.HealSkill(), robot2
		} else {
			return robot1.DodgeSkill(), robot2
		}
	}
}

func isExplosion(robot Robot) bool {
	if robot.heal <= 0 {
		return true
	}
	return false
}

func rollToDice(faceOfDice int) int8 {
	return int8(rand.Intn(faceOfDice) + 1)
}
