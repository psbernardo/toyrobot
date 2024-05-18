package pkg

import (
	"fmt"
	"testing"
)

func Expected(X int, Y int, faceDirection string) string {
	return fmt.Sprintf("%d,%d,%s", X, Y, faceDirection)
}

func Test_Robot_Case_1(t *testing.T) {
	toyRobot := NewRobot(East, 1, 2, NewMoveDirections(NewTableTop(4, 4)))
	toyRobot.Move()
	toyRobot.Move()
	toyRobot.TurnLeft()
	toyRobot.Move()

	executeTest(
		Expected(3, 3, North),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_2(t *testing.T) {
	toyRobot := NewRobot(North, 0, 0, NewMoveDirections(NewTableTop(5, 5)))
	toyRobot.TurnLeft()

	executeTest(
		Expected(0, 0, West),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_3(t *testing.T) {
	toyRobot := NewRobot(North, 0, 0, NewMoveDirections(NewTableTop(5, 5)))
	toyRobot.Move()

	executeTest(
		Expected(0, 1, North),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_4(t *testing.T) {
	toyRobot := NewRobot(East, 0, 0, NewMoveDirections(NewTableTop(4, 4)))
	toyRobot.Moves(5)
	toyRobot.TurnLeft()
	toyRobot.Move()

	executeTest(
		Expected(3, 1, North),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_5(t *testing.T) {
	toyRobot := NewRobot(North, 4, 4, NewMoveDirections(NewTableTop(4, 4)))
	toyRobot.Moves(2)
	toyRobot.TurnLeft()
	toyRobot.Moves(4)
	toyRobot.TurnLeft()
	toyRobot.Move()

	executeTest(
		Expected(0, 3, South),
		toyRobot.GetLocation(),
		t,
	)

}

func executeTest(expected, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("expected %s but got %s", expected, result)
	}
}
