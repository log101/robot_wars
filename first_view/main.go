package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

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
}

func (r Robot) Skill1(robot2 Robot) Robot {

	//robot2 = robot2.heal - int16(r.physicalPower)

	return robot2
}

var Deneme = "a"

func (r Robot) SetRobotProperty(name string, heal int16, speed uint8, physicalPower uint8, width uint8, length uint8, height uint8) Robot {
	r.Name = name
	r.heal = heal
	r.speed = speed
	r.physicalPower = physicalPower
	r.width = width
	r.length = length
	r.height = height
	r.volume = uint16(width) * uint16(length) * uint16(height)
	r.dodgePossibilty = uint8(10000/r.volume + uint16(speed))
	return r
}

func main() {
	var chosen int8
	var robots []Robot //robot struct'ından türeyecek
	robots = createRobot(robots, 10)
	//robots = createYourRobot(robots)
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
					fmt.Println(robots[len(robots)-1])
					chosen = -1
				}
			case 2:
				{
					printRobotList(robots)
					chosen = -1
				}
			case 4:
				{
					fmt.Println("Sıradaki rakip", robots[0])
				}
			case 5:
				{
					fmt.Println("Fight")
					chosen = -1
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

func createRobot(robots []Robot, countOfRobot int) []Robot {
	for i := 0; i < countOfRobot; i++ {
		//robotun değerleri ayarlanacak
		name := "robot" + strconv.Itoa(i)
		var heal int16 = int16(rand.Intn(21) + 80)          // heal
		var speed uint8 = uint8(rand.Intn(5) + 1)           // speed
		var physicalPower uint8 = uint8(rand.Intn(21) + 20) // physical power
		var width uint8 = uint8(rand.Intn(15) + 15)         // width
		var length uint8 = uint8(rand.Intn(15) + 15)        // length
		var height uint8 = uint8(rand.Intn(15) + 15)        // height
		var newRobot Robot
		newRobot = newRobot.SetRobotProperty(name, heal, speed, physicalPower, width, length, height)

		//en son robotlar eklenecek
		robots = append(robots, newRobot)
	}
	return robots
}

func printRobotList(robots []Robot) {
	for i := 0; i < len(robots); i++ {
		//isim güç falan yazılacak burada ekrana
		fmt.Println(robots[i].Name)
		fmt.Println(robots[i].heal)
		fmt.Println(robots[i].speed)
		fmt.Println(robots[i].physicalPower)
		fmt.Println(robots[i].volume)
		fmt.Println(robots[i].dodgePossibilty)
	}
}

func createYourRobot(robots []Robot) []Robot {
	var playerRobot Robot //robot struct'tından türeyecek
	fmt.Print("Robotunun ismi : ")
	fmt.Scanf("%s\n", &playerRobot.Name)
	//fmt.Scanf("%s\n", &playerRobot)
	robots = append(robots, playerRobot)
	return robots
}
