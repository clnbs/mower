package parser

import (
	"bufio"
	"errors"
	"jellysmack/internal/app/game"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FileParser struct {
	path string
}

func NewFileParser(path string) FileParser {
	return FileParser{path: path}
}

func (fp FileParser) Extract() (*game.Board, error) {
	lines, err := fp.extractLines()
	if err != nil {
		return nil, err
	}

	if len(lines)%2 != 1 || len(lines) < 3 {
		return nil, errors.New("invalid board definition")
	}

	rawBoardDefinition := lines[0]
	lines = lines[1:]
	boardSize, err := getPositionFromString(rawBoardDefinition)
	if err != nil {
		return nil, err
	}

	mowersAndMovements := []*game.MowerAndMovement{}

	for len(lines) != 0 {
		rawMowerDefinition := lines[0]
		mowerMovements := strings.Trim(lines[1], "\n")
		mowerPosition, err := getPositionFromString(rawMowerDefinition)
		if err != nil {
			return nil, err
		}
		mowerOrientation, err := getOrientationFromString(rawMowerDefinition)
		if err != nil {
			return nil, err
		}
		lines = lines[2:]

		mowersAndMovements = append(
			mowersAndMovements,
			&game.MowerAndMovement{
				Mower:     game.NewMower(mowerPosition[0], mowerPosition[1], mowerOrientation),
				Movements: mowerMovements,
			},
		)
	}

	board := game.NewBoard(boardSize[0], boardSize[1], mowersAndMovements)
	return board, nil
}

func (fp FileParser) extractLines() ([]string, error) {
	file, err := os.Open(fp.path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func getPositionFromString(input string) ([2]int, error) {
	input = strings.Trim(input, "\n")
	position := [2]int{}
	re := regexp.MustCompile(`[-]?\d[\d]*[\]?[\d{2}]*`)
	if !re.MatchString(input) {
		return [2]int{}, errors.New("unable to detect board definition")
	}
	submatch := re.FindAllString(input, -1)
	if len(submatch) != 2 {
		return [2]int{}, errors.New("file does not contains proper board definition")
	}
	position[0], _ = strconv.Atoi(submatch[0])
	position[1], _ = strconv.Atoi(submatch[1])
	return position, nil
}

func getOrientationFromString(input string) (rune, error) {
	input = strings.Trim(input, "\n")
	if len(input) == 0 {
		return 'Z', errors.New("empty input")
	}
	value := input[len(input)-1]
	return rune(value), nil
}
