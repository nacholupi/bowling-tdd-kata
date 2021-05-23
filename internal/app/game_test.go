package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestInitScoreZero(t *testing.T) {
	g := Game{}
	score := g.Score()

	assert.Equal(t, 0, score)
}
