package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_Score_InitZero(t *testing.T) {
	g := Game{}

	score := g.Score()

	assert.Equal(t, 0, score)
}

func TestGame_Roll(t *testing.T) {
	g := Game{}

	err := g.Roll(10)

	assert.NoError(t, err)
}

func TestGame_RollMoreThanTen(t *testing.T) {
	g := Game{}

	err := g.Roll(11)

	assert.Error(t, err)
}

func TestGame_RollLessThanZero(t *testing.T) {
	g := Game{}

	err := g.Roll(-1)

	assert.Error(t, err)
}
