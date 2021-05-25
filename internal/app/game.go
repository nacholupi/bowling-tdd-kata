package app

import "fmt"

type Game struct {
	score []int
}

const totalPins = 10
const totalFrames = 10

func (g Game) Score() int {

	totalScore := 0
	scoreIndex := 0

	for frameNum := 1; frameNum <= totalFrames; frameNum++ {

		totalScore += g.scoreValue(scoreIndex) + g.scoreValue(scoreIndex+1)

		if g.isStrike(scoreIndex) {
			totalScore += g.scoreValue(scoreIndex + 2)
			scoreIndex += 1
			continue
		}
		if g.isSpare(scoreIndex) {
			totalScore += g.scoreValue(scoreIndex + 2)
		}
		scoreIndex += 2
	}

	return totalScore
}

func (g Game) scoreValue(i int) int {
	if i >= len(g.score) {
		return 0
	}
	return g.score[i]
}

func (g Game) isSpare(i int) bool {
	return g.scoreValue(i)+g.scoreValue(i+1) == totalPins
}

func (g Game) isStrike(i int) bool {
	return g.scoreValue(i) == totalPins
}

func (g *Game) Roll(pins int) error {
	if pins > 10 || pins < 0 {
		return fmt.Errorf("illegal argument: choose a number between 0 and 10")
	}

	g.score = append(g.score, pins)
	return nil
}
