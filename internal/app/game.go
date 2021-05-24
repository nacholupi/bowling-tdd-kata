package app

import "fmt"

type Game struct {
	score []int
}

const totalPins = 10

func (g Game) Score() int {

	totalScore := 0

	bonusRolls := 0
	firstFrameShoot := true

	for i := range g.score {

		shootPoints := g.score[i]

		totalScore += shootPoints

		if bonusRolls > 0 {
			totalScore += shootPoints
			bonusRolls -= 1
		}

		if firstFrameShoot && shootPoints == totalPins {
			bonusRolls = 2
			continue
		}

		prevShootPoints := g.score[i-i]
		if !firstFrameShoot && shootPoints+prevShootPoints == totalPins {
			bonusRolls = 1
		}

		firstFrameShoot = !firstFrameShoot
	}

	return totalScore
}

func (g *Game) Roll(pins int) error {
	if pins > 10 || pins < 0 {
		return fmt.Errorf("illegal argument: choose a number between 0 and 10")
	}

	g.score = append(g.score, pins)
	return nil
}
