package game_test

import (
	"jellysmack/internal/app/game"
	"jellysmack/internal/pkg/orientation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMowerAndMovementMovement(t *testing.T) {
	mm := game.MowerAndMovement{
		Mower: &game.Mower{
			X:           0,
			Y:           0,
			Orientation: orientation.Orientation('N'),
		},
		Movements: "FFLFRF",
	}

	expected := []orientation.Orientation{
		orientation.Orientation('F'),
		orientation.Orientation('F'),
		orientation.Orientation('L'),
		orientation.Orientation('F'),
		orientation.Orientation('R'),
		orientation.Orientation('F'),
	}

	for _, e := range expected {
		move, err := mm.PopMovement()
		assert.Nil(t, err)
		assert.Equal(t, move, e)
	}

	assert.Equal(t, true, mm.Done())
}
