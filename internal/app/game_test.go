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

func TestGame_Roll_Rough_Patch(t *testing.T) {
	g := Game{}

	for i := 0; i < 20; i++ {
		_ = g.Roll(1)
	}

	assert.Equal(t, 20, g.Score())
}

func TestGame_Roll_First_Strike(t *testing.T) {
	g := Game{}

	_ = g.Roll(10)
	_ = g.Roll(1)
	_ = g.Roll(1)

	assert.Equal(t, 14, g.Score())
}

func TestGame_Roll_First_Spare(t *testing.T) {
	g := Game{}

	_ = g.Roll(5)
	_ = g.Roll(5)
	_ = g.Roll(2)
	_ = g.Roll(2)

	assert.Equal(t, 16, g.Score())
}

func TestGame_Roll_First_Spare_Starting_With_Zero(t *testing.T) {
	g := Game{}

	_ = g.Roll(0)
	_ = g.Roll(10)
	_ = g.Roll(2)
	_ = g.Roll(2)

	assert.Equal(t, 16, g.Score())
}
