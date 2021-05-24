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

func TestGame_Roll_Returns_Illegal_Arg(t *testing.T) {
	tests := []struct {
		name string
		pins int
	}{
		{
			name: "Less than zero",
			pins: -1,
		},
		{
			name: "More than ten",
			pins: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := Game{}

			err := g.Roll(test.pins)

			assert.Error(t, err)
		})
	}
}

func TestGame_Roll_Combinations(t *testing.T) {
	tests := []struct {
		name  string
		rolls []int
		score int
	}{
		{
			name: "Rough patch",
			rolls: []int{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			score: 20,
		},
		{
			name:  "First strike",
			rolls: []int{10, 1, 1},
			score: 14,
		},
		{
			name:  "First spare",
			rolls: []int{5, 5, 2, 2},
			score: 16,
		},
		{
			name:  "First spare starting with zero",
			rolls: []int{0, 10, 2, 2},
			score: 16,
		},
		{
			name:  "Two consecutive strikes",
			rolls: []int{10, 10, 1, 2},
			score: 37,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			g := Game{}

			for _, roll := range test.rolls {
				_ = g.Roll(roll)
			}

			assert.Equal(t, test.score, g.Score())
		})
	}
}
