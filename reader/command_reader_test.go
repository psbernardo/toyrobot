package reader

import (
	"fmt"
	"testing"

	"github.com/psbernardo/toyrobot/pkg"
)

func TestCommandReader(t *testing.T) {

	commandReader := NewCommandReader(GetDefeaultAction()...)
	robotLocation, err := commandReader.RunCommand(
		"PLACE 0,0,NORTH",
		MOVE,
	)

	checkRobotLocation(
		t,
		Expected(0, 1, pkg.North),
		robotLocation,
	)

	checkError(t, nil, err)

}

func TestCommandReader_With_Empty_Robot(t *testing.T) {

	command := NewCommandReader(GetDefeaultAction()...)
	robotLocation, err := command.RunCommand(
		MOVE,
		RIGHT,
		LEFT,
		"PLACE 1,2,EAST",
		MOVE,
		MOVE,
		LEFT,
		MOVE,
	)

	checkRobotLocation(
		t,
		Expected(3, 3, pkg.North),
		robotLocation,
	)

	checkError(t, nil, err)

}

func TestCommandReader_With_twice_Robot_Initialization(t *testing.T) {

	command := NewCommandReader(GetDefeaultAction()...)
	robotLocation, err := command.RunCommand(
		MOVE,
		RIGHT,
		LEFT,
		"PLACE 1,2,EAST",
		MOVE,
		MOVE,
		LEFT,
		MOVE,
		"PLACE 0,0,NORTH",
		LEFT,
	)

	checkRobotLocation(
		t,
		Expected(0, 0, pkg.West),
		robotLocation,
	)

	checkError(t, nil, err)

}

func TestCommandReader_Case_2(t *testing.T) {

	command := NewCommandReader(GetDefeaultAction()...)
	robotLocation, err := command.RunCommand(
		"PLACE 3,3,NORTH",
		MOVE,
		LEFT,
		MOVE,
		MOVE,
		MOVE,
	)

	checkRobotLocation(
		t,
		Expected(0, 4, pkg.West),
		robotLocation,
	)

	checkError(t, nil, err)

}

func TestCommandReader_Invalid_Place_Command(t *testing.T) {

	testCaseList := []struct {
		comands     []string
		expectedErr error
	}{
		{ // invalid Y coordinates
			comands: []string{
				LEFT,
				"PLACE 1,H,EAST",
			},
			expectedErr: ErrInvalidCoordinates("Y", "H"),
		},
		{ // invalid X coordinates
			comands: []string{
				"PLACE B1,4,EAST",
			},
			expectedErr: ErrInvalidCoordinates("X", "B1"),
		},
		{ // not enough place parameter
			comands: []string{
				"PLACE 1,4",
			},
			expectedErr: ErrNotEnoughParameter,
		},
		{ // not enough place parameter
			comands: []string{
				"PLACE",
			},
			expectedErr: ErrNotEnoughParameter,
		},
		{ // wrong face direction
			comands: []string{
				"PLACE 1,4,WWEST",
			},
			expectedErr: ErrInvalidDirection,
		},
		{ // coordinates is outside of the table size
			comands: []string{
				"PLACE 1,5,WEST",
			},
			expectedErr: ErrCoordinatesIsOutsideOfTableSize,
		},
		{ // invalid coordinates
			comands: []string{
				"PLACE 0,-1,WEST",
			},
			expectedErr: ErrInvalidLocationCoordinates,
		},
	}

	command := NewCommandReader(GetDefeaultAction()...)

	for _, testData := range testCaseList {
		robotLocation, err := command.RunCommand(testData.comands...)
		checkRobotLocation(t, "", robotLocation)
		checkError(t, testData.expectedErr, err)
	}

}

func Expected(X int, Y int, faceDirection string) string {
	return fmt.Sprintf("%d,%d,%s", X, Y, faceDirection)
}

func checkRobotLocation(t *testing.T, expectedLocation, robotLocation string) {

	if robotLocation != expectedLocation {
		t.Errorf("expected %s but got %s", expectedLocation, robotLocation)
	}

}

func checkError(t *testing.T, expectedError, commandErr error) {

	if checkErrNil(expectedError) != checkErrNil(commandErr) {
		t.Errorf("expected %s but got %s", checkErrNil(expectedError), checkErrNil(commandErr))
	}

}

func checkErrNil(err error) string {
	if err == nil {
		return "nil"
	}

	return err.Error()
}
