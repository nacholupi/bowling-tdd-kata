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

func TestGame_Roll_returns_illegal_arg(t *testing.T) {
	tests := []struct {
		name string
		pins int
	}{
		{
			name: "LESS_THAN_ZERO",
			pins: -1,
		},
		{
			name: "MORE_THAN_TEN",
			pins: 11,
		},
	}

	g := Game{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			err := g.Roll(test.pins)

			assert.Error(t, err)
		})
	}
}

func TestGame_Roll_retain_score(t *testing.T) {
	g := Game{}

	err := g.Roll(9)

	assert.NoError(t, err)
	assert.Equal(t, 9, g.Score())
}

func TestGame_Roll_add_score(t *testing.T) {
	g := Game{}

	_ = g.Roll(9)
	_ = g.Roll(5)

	assert.Equal(t, 14, g.Score())
}
