package pkg

import (
	"fmt"
	"testing"
)

func Expected(X int, Y int, faceDirection string) string {
	return fmt.Sprintf("%d,%d,%s", X, Y, faceDirection)
}

func Test_Robot_Case_1(t *testing.T) {
	tableTop := NewTableTop(5, 5)
	toyRobot := NewRobot(East, 1, 2)
	toyRobot.Move(tableTop)
	toyRobot.Move(tableTop)
	toyRobot.TurnLeft()
	toyRobot.Move(tableTop)

	executeTest(
		Expected(3, 3, North),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_2(t *testing.T) {

	toyRobot := NewRobot(North, 0, 0)
	toyRobot.TurnLeft()

	executeTest(
		Expected(0, 0, West),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_3(t *testing.T) {
	tableTop := NewTableTop(5, 5)

	toyRobot := NewRobot(North, 0, 0)
	toyRobot.Move(tableTop)

	executeTest(
		Expected(0, 1, North),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_4(t *testing.T) {
	tableTop := NewTableTop(4, 4)

	toyRobot := NewRobot(East, 0, 0)
	toyRobot.Moves(tableTop, 5)
	toyRobot.TurnLeft()
	toyRobot.Move(tableTop)

	executeTest(
		Expected(3, 1, North),
		toyRobot.GetLocation(),
		t,
	)

}

func Test_Robot_Case_5(t *testing.T) {
	tableTop := NewTableTop(4, 4)
	toyRobot := NewRobot(North, 4, 4)
	toyRobot.Moves(tableTop, 2)
	toyRobot.TurnLeft()
	toyRobot.Moves(tableTop, 4)
	toyRobot.TurnLeft()
	toyRobot.Move(tableTop)

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
