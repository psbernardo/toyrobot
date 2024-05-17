package pkg

import "fmt"

type ToyRobot struct {
	CurrentLocation *Location
	FaceDirection   string
	mapMoveFunc     map[string]MoveLocationFunc
}

// X= 1, Y = 2

// 4 |0,1,2,3,4
// 3 |0,1,2,3,4
// 2 |0,1,2,3,4
// 1 |0,1,2,3,4
// 0 |0,1,2,3,4

// it defends on direction we will add
// NORTH MOVE(y)++
// SOUTH MOVE(y)--

// EAST MOVE  (X)++)
// WEST MOVE  (X)--

// FACE- DIRECTION
// NORTH - L(WEST) |  R(EAST)
// SOUTH - L(EAST) |  R(WEST)
// EAST  - L(NORTH) | R(SOUTH)
// WEST  - L(SOUTH) | R (NORTH)

const (
	North = "NORTH"
	East  = "EAST"
	West  = "WEST"
	South = "SOUTH"
)

func IsValidDirection(value string) bool {
	for _, direction := range []string{
		North,
		East,
		West,
		South,
	} {
		if value == direction {
			return true
		}
	}

	return false
}

type MoveLocationFunc func(l *Location)

func NewRobot(faceDirection string, X, Y int, moveFunc map[string]MoveLocationFunc) *ToyRobot {
	return &ToyRobot{
		FaceDirection: faceDirection,
		CurrentLocation: &Location{
			X: X,
			Y: Y,
		},
		mapMoveFunc: moveFunc,
	}
}

func (r *ToyRobot) Move() {
	if moveFunc, OK := r.mapMoveFunc[r.FaceDirection]; OK {
		moveFunc(r.CurrentLocation)
	}
}

func (r *ToyRobot) Moves(count int) {
	for i := 0; i < count; i++ {
		r.Move()
	}
}

// FACE- DIRECTION
// NORTH - L(WEST) |  R(EAST)
// SOUTH - L(EAST) |  R(WEST)
// EAST  - L(NORTH) | R(SOUTH)
// WEST  - L(SOUTH) | R (NORTH)

func (r *ToyRobot) TurnRight() {
	switch r.FaceDirection {
	case North:
		r.FaceDirection = East
	case East:
		r.FaceDirection = South
	case West:
		r.FaceDirection = North
	case South:
		r.FaceDirection = West
	}
}

func (r *ToyRobot) TurnLeft() {
	switch r.FaceDirection {
	case North:
		r.FaceDirection = West
	case East:
		r.FaceDirection = North
	case West:
		r.FaceDirection = South
	case South:
		r.FaceDirection = East
	}

}

func (r ToyRobot) GetLocation() string {
	return fmt.Sprintf("%d,%d,%s", r.CurrentLocation.X, r.CurrentLocation.Y, r.FaceDirection)
}
