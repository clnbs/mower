package game

import (
	"errors"
	"jellysmack/internal/pkg/orientation"
)

type MowerAndMovement struct {
	Mower           *Mower
	Movements       string
	currentMovement int
}

func (mm *MowerAndMovement) PopMovement() (orientation.Orientation, error) {
	if mm.Done() {
		return 'Z', errors.New("no more movement available")
	}
	move := mm.Movements[mm.currentMovement]
	mm.currentMovement++
	return orientation.Orientation(move), nil
}

func (mm MowerAndMovement) Done() bool {
	if len(mm.Movements) <= mm.currentMovement {
		return true
	}
	return false
}
