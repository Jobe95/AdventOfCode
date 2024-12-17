package main

import (
	"advent-of-code-2024/utils"
	"regexp"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

type Robot struct {
	px, py, vx, vy int
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	seconds := 100
	rows, cols := 103, 101

	robots := []Robot{}

	for _, val := range input {
		matches := re.FindStringSubmatch(val)

		px, py := utils.ParseToInt(matches[1]), utils.ParseToInt(matches[2])
		vx, vy := utils.ParseToInt(matches[3]), utils.ParseToInt(matches[4])

		robot := Robot{px, py, vx, vy}
		robots = append(robots, robot)
	}

	for i := 0; i < seconds; i++ {
		for i := range robots {
			moveRobot(&robots[i], cols, rows)
		}
	}

	ySplit := rows / 2
	xSplit := cols / 2

	topLeft := []Robot{}
	topRight := []Robot{}
	bottomRight := []Robot{}
	bottomLeft := []Robot{}

	for i := range robots {
		robot := robots[i]
		if robot.px < xSplit && robot.py < ySplit {
			topLeft = append(topLeft, robot)
		} else if robot.px > xSplit && robot.py < ySplit {
			topRight = append(topRight, robot)
		} else if robot.px > xSplit && robot.py > ySplit {
			bottomRight = append(bottomRight, robot)
		} else if robot.px < xSplit && robot.py > ySplit {
			bottomLeft = append(bottomLeft, robot)
		}
	}

	output = len(topLeft) * len(topRight) * len(bottomRight) * len(bottomLeft)

	return output
}

func partTwo() int {
	output := 0
	return output
}

func moveRobot(robot *Robot, cols, rows int) {
	newX, newY := (cols+robot.px+robot.vx)%cols, (rows+robot.py+robot.vy)%rows
	// fmt.Printf("OX=%d, OY=%d, NX=%d, NY=%d", robot.px, robot.py, newX, newY)
	robot.px = newX
	robot.py = newY
}
