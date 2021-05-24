package app

import "fmt"

type Game struct {
	score []int
}

func (g Game) Score() int {
	totalScore := 0

	bonusRolls := 0
	firstFrameShoot := true

	for i := range g.score {

		shootScore := g.score[i]

		totalScore += shootScore

		if firstFrameShoot && shootScore == 10 {
			bonusRolls = 2
			continue
		}

		if bonusRolls > 0 {
			totalScore += shootScore
			bonusRolls -= 1
		}

		prevShootScore := g.score[i-i]
		if !firstFrameShoot && shootScore+prevShootScore == 10 {
			bonusRolls = 1
		}

		if firstFrameShoot == true {
			firstFrameShoot = false
		} else {
			firstFrameShoot = true
		}
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
