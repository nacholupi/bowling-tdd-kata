package app

import "fmt"

type Game struct {
	score []int
}

func (g Game) Score() int {
	totalScore := 0
	for _, s := range g.score {
		totalScore += s
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
