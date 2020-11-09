package worldmap

import (
	"fmt"
)

// Direction type for collision checking.
type Direction int

// Enum for directions.
const (
	Up Direction = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
)

// GetDirection converts an integer into a Direction.
func GetDirection(i int) (Direction, error) {
	switch i {
	case 0:
		return Up, nil
	case 1:
		return Down, nil
	case 2:
		return Left, nil
	case 3:
		return Right, nil
	}

	return Up, fmt.Errorf("no such direction: %d", i)
}

// DirectionToText takes the integer
// representation for a direction and
// returns the name of the direction.
func DirectionToText(dir Direction) string {
	switch dir {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}

	return "Unknown direction"
}

// IndexAtDirection returns the index
// at the direction from the current index.
func IndexAtDirection(idx int, dir Direction) int {
	switch dir {

	case Up:
		return OneTileUp(idx)
	case Right:
		return OneTileRight(idx)
	case Down:
		return OneTileDown(idx)
	case Left:
		return OneTileLeft(idx)
	case UpLeft:
		return OneTileUpLeft(idx)
	case UpRight:
		return OneTileUpRight(idx)
	case DownLeft:
		return OneTileDownLeft(idx)
	case DownRight:
		return OneTileDownRight(idx)

	default:
		fmt.Println("unknown direction:", DirectionToText(dir))
		return -1
	}
}

// OneTileUp returns idx one row up.
func OneTileUp(idx int) int {
	return idx - TilesW
}

// OneTileDown returns idx one row down.
func OneTileDown(idx int) int {
	return idx + TilesW
}

// OneTileLeft returns idx one column left.
func OneTileLeft(idx int) int {
	return idx - 1
}

// OneTileRight returns idx one column right.
func OneTileRight(idx int) int {
	return idx + 1
}

// OneTileUpLeft returns idx one column left,
// one row up.
func OneTileUpLeft(idx int) int {
	return OneTileUp(OneTileLeft(idx))
}

// OneTileUpRight returns idx one column right,
// one row up.
func OneTileUpRight(idx int) int {
	return OneTileUp(OneTileRight(idx))
}

// OneTileDownLeft returns idx one column left,
// one row up.
func OneTileDownLeft(idx int) int {
	return OneTileDown(OneTileLeft(idx))
}

// OneTileDownRight returns idx one column right,
// one row down.
func OneTileDownRight(idx int) int {
	return OneTileDown(OneTileRight(idx))
}

// NextIdxToDir returns the direction needed
// to traverse to the next position
func NextIdxToDir(idx, next int) (Direction, error) {
	if next == OneTileUp(idx) {
		return Up, nil
	}

	if next == OneTileDown(idx) {
		return Down, nil
	}

	if next == OneTileLeft(idx) {
		return Left, nil
	}

	if next == OneTileRight(idx) {
		return Right, nil
	}

	return Up, fmt.Errorf("idx and next not adjacent")
}
