package game

import (
	"fmt"
	"jellysmack/internal/pkg/orientation"
)

type Mower struct {
	X           int
	Y           int
	Orientation orientation.Orientation
}

func NewMower(x, y int, o rune) *Mower {
	return &Mower{
		X:           x,
		Y:           y,
		Orientation: orientation.Orientation(o),
	}
}

func (m Mower) String() string {
	return fmt.Sprintf("%d %d %c", m.X, m.Y, m.Orientation)
}

func (m Mower) Position() [2]int {
	return [2]int{m.X, m.Y}
}

func (m *Mower) SimulateMove() [2]int {
	return m.computeMove()
}

func (m *Mower) MoveForward() {
	newPosition := m.computeMove()
	m.X = newPosition[0]
	m.Y = newPosition[1]
}

func (m *Mower) Rotate(direction orientation.Orientation) {
	switch direction {
	case 'L':
		m.Orientation = m.Orientation.Left()
	case 'R':
		m.Orientation = m.Orientation.Right()
	}
}

func (m *Mower) computeMove() [2]int {
	position := [2]int{m.X, m.Y}
	switch m.Orientation {
	case 'N':
		position[1] += 1
	case 'E':
		position[0] += 1
	case 'W':
		position[0] -= 1
	case 'S':
		position[1] -= 1
	}
	return position
}
