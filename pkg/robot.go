package pkg

import "fmt"

type ToyRobot struct {
	CurrentLocation *Location
	FaceDirection   string
	mapMoveFunc     map[string]moveLocationFunc
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

type moveLocationFunc func(tabletop TableTop)

func NewRobot(faceDirection string, X, Y int) *ToyRobot {

	toyRobot := &ToyRobot{
		FaceDirection: faceDirection,
		CurrentLocation: &Location{
			X: X,
			Y: Y,
		},
	}

	toyRobot.mapMoveFunc = map[string]moveLocationFunc{
		North: toyRobot.moveNorth,
		East:  toyRobot.moveEast,
		West:  toyRobot.moveWest,
		South: toyRobot.moveSouth,
	}

	return toyRobot

}

func (r *ToyRobot) Move(tabletop TableTop) {
	if moveFunc, OK := r.mapMoveFunc[r.FaceDirection]; OK {
		moveFunc(tabletop)
	}
}

func (r *ToyRobot) Moves(tabletop TableTop, count int) {
	for i := 0; i < count; i++ {
		r.Move(tabletop)
	}
}

func (r *ToyRobot) moveNorth(tabletop TableTop) {
	if r.CurrentLocation.Y < (tabletop.Y - 1) {
		r.CurrentLocation.Y++
	}
}

func (r *ToyRobot) moveSouth(tabletop TableTop) {
	if r.CurrentLocation.Y > 0 {
		r.CurrentLocation.Y--
	}

}

func (r *ToyRobot) moveEast(tabletop TableTop) {
	if r.CurrentLocation.X < (tabletop.X - 1) {
		r.CurrentLocation.X++
	}
}

func (r *ToyRobot) moveWest(tabletop TableTop) {

	if r.CurrentLocation.X > 0 {
		r.CurrentLocation.X--
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
