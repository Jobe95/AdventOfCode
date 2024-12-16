package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

type Button struct {
	TokenCost int
	X         int
	Y         int
}

type Claw struct {
	A      Button
	B      Button
	PrizeX int
	PrizeY int
}

func (b Button) String() string {
	return fmt.Sprintf("Button(TokenCost=%d, X=%d, Y=%d)", b.TokenCost, b.X, b.Y)
}

func (c Claw) String() string {
	return fmt.Sprintf("Claw(A=%s, B=%s, PrizeX=%d, PrizeY=%d)", c.A, c.B, c.PrizeX, c.PrizeY)
}

func partOne() int {
	output := 0
	input := utils.ReadFile("input.txt")
	groups := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	claws := createClaws(groups)

	for i := 0; i < len(claws); i++ {
		claw := claws[i]

		a, b := solveClaw(claw)
		output += a * claw.A.TokenCost
		output += b * claw.B.TokenCost
	}

	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFile("input.txt")
	groups := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	claws := createClaws(groups)

	for i := 0; i < len(claws); i++ {
		claw := claws[i]
		claw.PrizeY = 10000000000000 + claw.PrizeY
		claw.PrizeX = 10000000000000 + claw.PrizeX
		a, b := solveClaw(claw)
		output += a * claw.A.TokenCost
		output += b * claw.B.TokenCost
	}

	return output
}

func solveClaw(c Claw) (int, int) {
	d := c.A.X*c.B.Y - c.B.X*c.A.Y
	d1 := c.PrizeX*c.B.Y - c.PrizeY*c.B.X
	d2 := c.PrizeY*c.A.X - c.PrizeX*c.A.Y

	if d1%d != 0 || d2%d != 0 {
		return 0, 0
	}

	return d1 / d, d2 / d
}

func parseXY(line string) (int, int) {
	re := regexp.MustCompile(`X[+=](\d+), Y[+=](\d+)`)
	matches := re.FindStringSubmatch(line)
	x, y := utils.ParseToInt(matches[1]), utils.ParseToInt(matches[2])

	return x, y
}

func createClaws(groups []string) []Claw {
	claws := []Claw{}
	for _, group := range groups {
		lines := strings.Split(group, "\n")
		if len(lines) != 3 {
			panic("Out of bounds!")
		}

		buttonAX, buttonAY := parseXY(lines[0])
		buttonBX, buttonBY := parseXY(lines[1])
		prizeX, prizeY := parseXY(lines[2])

		claw := Claw{
			A: Button{
				TokenCost: 3,
				X:         buttonAX,
				Y:         buttonAY,
			},
			B: Button{
				TokenCost: 1,
				X:         buttonBX,
				Y:         buttonBY,
			},
			PrizeX: prizeX,
			PrizeY: prizeY,
		}
		claws = append(claws, claw)
	}
	return claws
}

// func gcd(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	return a
// }
