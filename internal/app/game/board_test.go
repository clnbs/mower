package game_test

import (
	"jellysmack/internal/app/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardIsMovePossible(t *testing.T) {
	mm := []*game.MowerAndMovement{
		{
			Mower:     game.NewMower(0, 0, 'N'),
			Movements: "RFFFF",
		},
		{
			Mower:     game.NewMower(2, 1, 'N'),
			Movements: "RFFFF",
		},
	}
	b := game.NewBoard(10, 10, mm)

	testCases := []struct {
		Input    [2]int
		Expected bool
		Name     string
	}{
		{
			Input:    [2]int{1, 2},
			Expected: true,
			Name:     "test case 1",
		},
		{
			Input:    [2]int{2, 1},
			Expected: false,
			Name:     "test case 2",
		},
		{
			Input:    [2]int{0, 0},
			Expected: false,
			Name:     "test case 3",
		},
	}

	for _, c := range testCases {
		result := b.IsMovePossible(c.Input)
		assert.Equal(t, c.Expected, result, c.Name)
	}
}
