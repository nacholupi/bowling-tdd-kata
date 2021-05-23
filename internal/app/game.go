package app

import "fmt"

type Game struct {
	score int
}

func (g Game) Score() int {
	return g.score
}

func (g Game) Roll(pins int) error {
	if pins > 10 || pins < 0 {
		return fmt.Errorf("illegal argument: choose a number between 0 and 10")
	}
	return nil
}
