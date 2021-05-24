package app

import "fmt"

type Game struct {
	score []int
}

const totalPins = 10

func (g Game) Score() int {

	totalScore := 0

	firstFrameShoot := true

	for i := range g.score {

		shootPoints := g.score[i]

		totalScore += shootPoints

		if g.isStrike(firstFrameShoot, i) {
			totalScore += g.score[i+1] + g.score[i+2]
			continue
		}

		if g.isSpare(firstFrameShoot, i) {
			totalScore += g.score[i+1]
		}

		firstFrameShoot = !firstFrameShoot
	}

	return totalScore
}

func (g Game) isSpare(firstFrameShoot bool, i int) bool {
	return !firstFrameShoot && g.score[i]+g.score[i-1] == totalPins
}

func (g Game) isStrike(firstFrameShoot bool, i int) bool {
	return firstFrameShoot && g.score[i] == totalPins
}

func (g *Game) Roll(pins int) error {
	if pins > 10 || pins < 0 {
		return fmt.Errorf("illegal argument: choose a number between 0 and 10")
	}

	g.score = append(g.score, pins)
	return nil
}
