package parser

import "jellysmack/internal/app/game"

type Parser interface {
	Extract() (*game.Board, error)
}
