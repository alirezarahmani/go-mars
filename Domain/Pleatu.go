package Domain

import (
	"fmt"
	"reflect"
)

type Plateau struct {
	X      int
	Y      int
	rovers []MarsRover
}

// NewPlateau creates a new valid name object
func NewPlateau(x int, y int) (Plateau, error) {
	if x < 0 || y < 0  {
		return Plateau{}, fmt.Errorf("invalid position")
	}
	return Plateau{X: x, Y: y}, nil
}

func (p Plateau) addRover(marsRover *MarsRover)  {
	AppendIfMissing(p.rovers, marsRover)
}

func (p Plateau) isOccupied(position Position) bool {
	for _, v := range p.rovers {
		x , y := v.position.String()
		x1, y1 := position.String()
		if x == x1 && y == y1 {
			return true
		}
	}
	return false
}

func AppendIfMissing(slice []MarsRover, i *MarsRover) []MarsRover {
	for _, ele := range slice {
		if reflect.DeepEqual(ele, i) {
			return slice
		}
	}
	return append(slice, *i)
}