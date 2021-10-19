package game_test

import (
	"jellysmack/internal/app/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	testCase := struct {
		Input    [2]int
		Expected [2]int
		Name     string
	}{
		Input:    [2]int{2, 3},
		Expected: [2]int{2, 3},
		Name:     "basic test",
	}

	mow := game.NewMower(testCase.Input[0], testCase.Input[1], 'N')
	result := mow.Position()
	assert.Equal(t, testCase.Expected, result, testCase.Name)
}

func TestMoveForwardNorth(t *testing.T) {
	mow := game.NewMower(5, 5, 'N')
	testCases := []struct {
		Expected [2]int
		Name     string
	}{
		{
			Expected: [2]int{5, 6},
			Name:     "move forward north x1",
		},
		{
			Expected: [2]int{5, 7},
			Name:     "move forward north x2",
		},
		{
			Expected: [2]int{5, 8},
			Name:     "move forward north x3",
		},
		{
			Expected: [2]int{5, 9},
			Name:     "move forward north x4",
		},
	}

	for _, c := range testCases {
		mow.MoveForward()
		assert.Equal(t, c.Expected, mow.Position(), c.Name)
	}
}

func TestMoveForwardSouth(t *testing.T) {
	mow := game.NewMower(5, 5, 'S')
	testCases := []struct {
		Expected [2]int
		Name     string
	}{
		{
			Expected: [2]int{5, 4},
			Name:     "move forward south x1",
		},
		{
			Expected: [2]int{5, 3},
			Name:     "move forward south x2",
		},
		{
			Expected: [2]int{5, 2},
			Name:     "move forward south x3",
		},
		{
			Expected: [2]int{5, 1},
			Name:     "move forward south x4",
		},
	}

	for _, c := range testCases {
		mow.MoveForward()
		assert.Equal(t, c.Expected, mow.Position(), c.Name)
	}
}

func TestMoveForwardEast(t *testing.T) {
	mow := game.NewMower(5, 5, 'E')
	testCases := []struct {
		Expected [2]int
		Name     string
	}{
		{
			Expected: [2]int{6, 5},
			Name:     "move forward east x1",
		},
		{
			Expected: [2]int{7, 5},
			Name:     "move forward east x2",
		},
		{
			Expected: [2]int{8, 5},
			Name:     "move forward east x3",
		},
		{
			Expected: [2]int{9, 5},
			Name:     "move forward east x4",
		},
	}

	for _, c := range testCases {
		mow.MoveForward()
		assert.Equal(t, c.Expected, mow.Position(), c.Name)
	}
}

func TestMoveForwardWest(t *testing.T) {
	mow := game.NewMower(5, 5, 'W')
	testCases := []struct {
		Expected [2]int
		Name     string
	}{
		{
			Expected: [2]int{4, 5},
			Name:     "move forward east x1",
		},
		{
			Expected: [2]int{3, 5},
			Name:     "move forward east x2",
		},
		{
			Expected: [2]int{2, 5},
			Name:     "move forward east x3",
		},
		{
			Expected: [2]int{1, 5},
			Name:     "move forward east x4",
		},
	}

	for _, c := range testCases {
		mow.MoveForward()
		assert.Equal(t, c.Expected, mow.Position(), c.Name)
	}
}
