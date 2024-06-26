package reader

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/psbernardo/toyrobot/pkg"
)

var (
	ErrNotEnoughParameter              = errors.New("command has not enough parameter")
	ErrInvalidDirection                = errors.New("invalid direction")
	ErrCoordinatesIsOutsideOfTableSize = errors.New("coordinates is outside of the table size")
	ErrInvalidLocationCoordinates      = errors.New("invalid coordinates")
)

func ErrInvalidCoordinates(coordinate, value string) error {
	return fmt.Errorf("invalid %s coordinates value[%s]", coordinate, value)
}

func Place(c *CommandReader) ActionHandler {

	// Create Robot from string
	createToyRobotFromString := func(command string) (*pkg.ToyRobot, error) {
		Xstring, Ystring, faceDirection, err := SplitByThree(command, ",")
		if err != nil {
			return nil, err
		}

		if !pkg.IsValidDirection(faceDirection) {
			return nil, ErrInvalidDirection
		}

		X, Y, err := CreateXYFromString(Xstring, Ystring)
		if err != nil {
			return nil, err
		}

		// coordinates is outside of the tablepot size
		if !c.TableTop.IsInsideTheTablePot(X, Y) {
			return nil, ErrCoordinatesIsOutsideOfTableSize
		}

		return pkg.NewRobot(faceDirection, X, Y), nil
	}

	return ActionHandler{
		Code:        PLACE,
		Description: "PLACE X,Y,F - will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.",
		Run: func(command string) error {

			index := strings.Index(command, " ")
			if index == -1 {
				return ErrNotEnoughParameter
			}

			ToyRobot, err := createToyRobotFromString(command[index:])
			if err != nil {

				return err
			}
			c.ToyRobot = ToyRobot
			return nil
		},
	}

}

func Move(c *CommandReader) ActionHandler {

	return ActionHandler{
		Code:        MOVE,
		Description: "MOVE   - will move the toy robot one unit forward in the direction it is currently facing.",
		Run: func(command string) error {
			if c.ToyRobot != nil {
				c.ToyRobot.Move(c.TableTop)
			}
			return nil
		},
	}

}

func Left(c *CommandReader) ActionHandler {

	return ActionHandler{
		Code:        LEFT,
		Description: "LEFT   - will rotate the robot 90 degrees in the specified direction without changing the position of the robot.",
		Run: func(command string) error {
			if c.ToyRobot != nil {
				c.ToyRobot.TurnLeft()
			}
			return nil
		},
	}

}

func Right(c *CommandReader) ActionHandler {

	return ActionHandler{
		Code:        RIGHT,
		Description: "RIGHT  - will rotate the robot 90 degrees in the specified direction without changing the position of the robot.",
		Run: func(command string) error {
			if c.ToyRobot != nil {
				c.ToyRobot.TurnRight()
			}
			return nil
		},
	}

}

func Report(c *CommandReader) ActionHandler {
	return ActionHandler{
		Code:        REPORT,
		Description: "REPORT - will announce the X,Y and F of the robot. This can be in any form, but standard output is sufficient. ",
		Run: func(command string) error {
			if c.ToyRobot != nil {
				fmt.Printf("Output: %s \n", c.ToyRobot.GetLocation())
			}
			return nil
		},
	}
}

func Help(c *CommandReader) ActionHandler {
	return ActionHandler{
		Code: HELP,
		Run: func(command string) error {

			for _, actionHandler := range c.mapCommandExecuter {
				if len(actionHandler.Description) > 0 {
					typeLines(actionHandler.Description)
				}

			}
			return nil
		},
	}
}

func Example(c *CommandReader) ActionHandler {
	return ActionHandler{
		Code: "EXAMPLE",
		Run: func(command string) error {

			exampleList := [][]string{
				{
					"a) PLACE 0,0,NORTH",
					"   MOVE",
					"   REPORT",
					"   Output: 0,1,NORTH\n",
				},
				{
					"b) PLACE 0,0,NORTH",
					"   LEFT",
					"   REPORT",
					"   Output: 0,0,WEST\n",
				},
				{
					"c) PLACE 1,2,EAST",
					"   MOVE",
					"   MOVE",
					"   LEFT",
					"   MOVE",
					"   REPORT",
					"   Output: 3,3,NORTH\n",
				},
			}

			for _, example := range exampleList {
				typeLines(example...)

			}
			return nil
		},
	}
}

func CreateXYFromString(x, y string) (int, int, error) {
	X, err := strconv.Atoi(x)
	if err != nil {
		return 0, 0, ErrInvalidCoordinates("X", x)
	}

	if X < 0 {
		return 0, 0, ErrInvalidLocationCoordinates
	}

	Y, err := strconv.Atoi(y)
	if err != nil {
		return 0, 0, ErrInvalidCoordinates("Y", y)
	}

	if Y < 0 {
		return 0, 0, ErrInvalidLocationCoordinates
	}

	return X, Y, nil
}

func typeLines(lines ...string) {
	for _, line := range lines {
		letters := []rune(line)
		for _, letter := range letters {
			fmt.Print(string(letter))
			n := rand.Intn(10)
			time.Sleep(time.Duration(n) * time.Nanosecond)
		}
		fmt.Println()
	}

}
