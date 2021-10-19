package orientation

import "errors"

var orientations = []Orientation{'N', 'E', 'S', 'W'}

type Orientation rune

func NewOrientation(o rune) (Orientation, error) {
	usable := false
	for _, possible := range orientations {
		if possible == Orientation(o) {
			usable = true
		}
	}
	if !usable {
		return orientations[0], errors.New("orientable not allowed")
	}
	return Orientation(o), nil
}

func (o Orientation) Left() Orientation {
	for index, orient := range orientations {
		if orient == o {
			if index == 0 {
				return orientations[len(orientations)-1]
			}
			return orientations[index-1]
		}
	}
	return o
}

func (o Orientation) Right() Orientation {
	for index, orient := range orientations {
		if orient == o {
			if index == len(orientations)-1 {
				return orientations[0]
			}
			return orientations[index+1]
		}
	}
	return o
}
