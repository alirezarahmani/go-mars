package Domain

import (
	"errors"
)

type MarsRover struct {
	position Position
	heading  DirectionType
	plateau  Plateau
}

// NewMarsRover class
func NewMarsRover(plateau Plateau, position Position, directionType DirectionType) (*MarsRover, error) {
	if position.IsOnPlateau(plateau) == false {
		return nil, errors.New("out of plateau")
	}
	if plateau.isOccupied(position) {
		return nil, errors.New("already Occupied")
	}
	mars := &MarsRover{
		position: position,
		plateau:  plateau,
		heading:  directionType,
	}
	plateau.addRover(mars)

	return mars, nil
}

func (mars *MarsRover) Move(move RoverMoveType) error {
	var err error
	switch move {
	case RoverMove.L:
		mars.turnLeft()
	case RoverMove.R:
		mars.turnRight()
	case RoverMove.M:
		err = mars.moveForward()
	}
	return err
}

func (mars *MarsRover) turnLeft() {
	switch mars.heading {
	case Direction.N:
		mars.heading = Direction.W
	case Direction.W:
		mars.heading = Direction.S
	case Direction.E:
		mars.heading = Direction.N
	case Direction.S:
		mars.heading = Direction.E
	}
}

func (mars *MarsRover) ReportPosition() (int, int) {
	return mars.position.X, mars.position.Y
}

func (mars *MarsRover) ReportHeading() DirectionType {
	return mars.heading
}

func (mars *MarsRover) turnRight() {
	switch mars.heading {
	case Direction.N:
		mars.heading = Direction.E
	case Direction.W:
		mars.heading = Direction.N
	case Direction.E:
		mars.heading = Direction.S
	case Direction.S:
		mars.heading = Direction.W
	}
}

func (mars *MarsRover) moveForward() error {
	var err error
	switch mars.heading {
	case Direction.N:
		mars.position, err = NewPosition(mars.position.X, mars.position.Y+1)
	case Direction.W:
		mars.position, err = NewPosition(mars.position.X, mars.position.Y-1)
	case Direction.E:
		mars.position, err = NewPosition(mars.position.X+1, mars.position.Y)
	case Direction.S:
		mars.position, err = NewPosition(mars.position.X-1, mars.position.Y)
	}
	return err
}
