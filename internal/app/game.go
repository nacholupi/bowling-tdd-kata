package app

type Game struct {
	score int
}

func (g Game) Score() int {
	return g.score
}
