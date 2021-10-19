package orientation_test

import (
	"jellysmack/internal/pkg/orientation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	testCases := []struct {
		Input    orientation.Orientation
		Expected orientation.Orientation
		Name     string
	}{
		{
			Input:    'N',
			Expected: 'W',
			Name:     "left from north",
		},
		{
			Input:    'E',
			Expected: 'N',
			Name:     "left from east",
		},
		{
			Input:    'S',
			Expected: 'E',
			Name:     "left from south",
		},
		{
			Input:    'W',
			Expected: 'S',
			Name:     "left from west",
		},
	}
	for _, c := range testCases {
		result := c.Input.Left()
		assert.Equal(t, string(c.Expected), string(result), c.Name)
	}
}

func TestRight(t *testing.T) {
	testCases := []struct {
		Input    orientation.Orientation
		Expected orientation.Orientation
		Name     string
	}{
		{
			Input:    'N',
			Expected: 'E',
			Name:     "right from north",
		},
		{
			Input:    'E',
			Expected: 'S',
			Name:     "right from east",
		},
		{
			Input:    'S',
			Expected: 'W',
			Name:     "right from south",
		},
		{
			Input:    'W',
			Expected: 'N',
			Name:     "right from west",
		},
	}
	for _, c := range testCases {
		result := c.Input.Right()
		assert.Equal(t, string(result), string(c.Expected), c.Name)
	}
}
