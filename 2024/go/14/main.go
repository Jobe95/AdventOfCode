package main

import (
	"advent-of-code-2024/utils"
	"fmt"
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
		for j := range robots {
			moveRobot(&robots[j], cols, rows)
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
	input := utils.ReadFileLinesAsStringArray("input.txt")

	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	seconds := 10000
	rows, cols := 103, 101

	robots := []Robot{}

	for _, val := range input {
		matches := re.FindStringSubmatch(val)

		px, py := utils.ParseToInt(matches[1]), utils.ParseToInt(matches[2])
		vx, vy := utils.ParseToInt(matches[3]), utils.ParseToInt(matches[4])

		robot := Robot{px, py, vx, vy}
		robots = append(robots, robot)
	}

	// for i := 0; i < seconds; i++ {
	// 	for j := range robots {
	// 		moveRobot(&robots[j], cols, rows)
	// 	}
	// }

	for i := 1; i < seconds; i++ {
		grid := utils.CreateGridFromDimensions(rows, cols, func(y, x int) string {
			return " "
		})
		for j := range robots {
			moveRobot(&robots[j], cols, rows)
			grid.Set(robots[j].py, robots[j].px, "#")
		}

		for _, row := range grid.Cells {
			count := 0

			for _, cell := range row {
				if cell.Item == "#" {
					count++
				}
			}

			if count > 30 {
				fmt.Println()
				fmt.Println(i)
				grid.PrintGridRaw()
				fmt.Println()
			}

		}
	}

	return output
}

func moveRobot(robot *Robot, cols, rows int) {
	newX, newY := (cols+robot.px+robot.vx)%cols, (rows+robot.py+robot.vy)%rows
	robot.px = newX
	robot.py = newY
	// fmt.Printf("OX=%d, OY=%d, NX=%d, NY=%d", robot.px, robot.py, newX, newY)
}
