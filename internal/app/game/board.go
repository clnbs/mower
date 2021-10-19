package game

import (
	"fmt"
	"log"
	"sync"
)

type Board struct {
	Mowers []*MowerAndMovement
	Size   [2]int
	wg     *sync.WaitGroup
}

func NewBoard(x, y int, mowers []*MowerAndMovement) *Board {
	fmt.Println("board size :", x, y)
	return &Board{
		Mowers: mowers,
		Size:   [2]int{x, y},
		wg:     &sync.WaitGroup{},
	}
}

func (b *Board) Play() {
	for !b.isDone() {
		b.step()
	}
	for _, m := range b.Mowers {
		fmt.Println(m.Mower.String())
	}
}

func (b Board) IsMovePossible(currentPosition [2]int) bool {
	if currentPosition[0] > b.Size[0] || currentPosition[0] < 0 ||
		currentPosition[1] > b.Size[1] || currentPosition[1] < 0 {
		return false
	}
	for _, m := range b.Mowers {
		if m.Mower.X == currentPosition[0] && m.Mower.Y == currentPosition[1] {
			return false
		}
	}
	return true
}

func (b *Board) step() {
	for _, m := range b.Mowers {
		b.wg.Add(1)
		go b.stepRoutine(m)
	}
	b.wg.Wait()
}

func (b *Board) stepRoutine(m *MowerAndMovement) {
	defer b.wg.Done()
	if m.Done() {
		return
	}

	move, err := m.PopMovement()
	if err != nil {
		log.Printf("mower can not move : %s", err)
	}

	if move != 'F' {
		m.Mower.Rotate(move)
		return
	}

	if b.IsMovePossible(m.Mower.SimulateMove()) {
		m.Mower.MoveForward()
	}
}

func (b Board) isDone() bool {
	for _, m := range b.Mowers {
		if !m.Done() {
			return false
		}
	}
	return true
}
