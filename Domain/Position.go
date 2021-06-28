package Domain

import (
	"encoding/json"
	"fmt"
)

type Position struct {
	// explicitly not public field
	X int
	Y int
}

// NewPosition creates a new valid name object
func NewPosition(x int, y int) (Position, error) {
	if x < 0 || y < 0 {
		return Position{}, fmt.Errorf("invalid position")
	}
	return Position{X: x, Y: y}, nil
}

func (p Position) IsOnPlateau(pl Plateau) bool {
	if p.X < 0 || p.X > pl.X {
		return false
	}
	if p.Y < 0 || p.Y > pl.Y {
		return false
	}
	return true
}

// String implements the fmt.Stringer interface.
func (p Position) String() (int, int) {
	return p.X, p.Y
}

// MarshalText used to serialize the object
func (p Position) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf(`{X: %d, Y: %d}`, p.X, p.Y)), nil
}

// UnmarshalText deserializes the object and returns an error if it's invalid.
func (p *Position) UnmarshalText(data []byte) error {
	var err error
	err = json.Unmarshal(data, &p)
	*p, err = NewPosition(p.X, p.Y)
	return err
}
