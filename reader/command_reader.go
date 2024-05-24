package reader

import (
	"strings"

	"github.com/psbernardo/toyrobot/pkg"
)

// COMMAND
const (
	PLACE  = "PLACE"
	MOVE   = "MOVE"
	LEFT   = "LEFT"
	RIGHT  = "RIGHT"
	REPORT = "REPORT"
	HELP   = "HELP"
)

type ActionHandler struct {
	Code        string
	Description string
	Run         func(command string) error
}

type CommandReaderOption func(r *CommandReader) ActionHandler

type CommandReader struct {
	ToyRobot           *pkg.ToyRobot
	TableTop           pkg.TableTop
	mapCommandExecuter map[string]ActionHandler
}

func NewCommandReader(options ...CommandReaderOption) *CommandReader {
	c := &CommandReader{
		mapCommandExecuter: make(map[string]ActionHandler),
		TableTop:           pkg.NewTableTop(5, 5),
	}
	for _, option := range options {
		actionHandler := option(c)
		c.mapCommandExecuter[actionHandler.Code] = actionHandler
	}
	return c
}

func (c *CommandReader) RunCommand(commands ...string) (string, error) {
	for _, command := range commands {

		if len(command) == 0 {
			continue
		}

		if actionHandler, OK := c.mapCommandExecuter[commandDecoder(command)]; OK {
			if err := actionHandler.Run(command); err != nil {
				return "", err
			}
		}
	}

	if c.ToyRobot != nil {
		return c.ToyRobot.GetLocation(), nil
	}

	return "", nil
}

func PrintInstruction() {
	typeLines([]string{
		"* This application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units",
		"* The robot is free to roam around the surface of the table, but must be prevented from falling to destruction",
		"* The first valid command to the robot is a PLACE command, after that, any sequence of commands may be issued, in any order, including another PLACE command.",
		"  The application should discard all commands in the sequence until a valid PLACE command has been executed.",
		"* A robot that is not on the table can choose the ignore the MOVE, LEFT, RIGHT and REPORT commands.",
		"* Any movement that would result in the robot falling from the table must be prevented, however further valid movement commands must still be allowed",
		"Type \"HELP\" to show the commands and \"EXAMPLE\" to show sample  Input/Output",
	}...)
}

// always get the first word from the string to get the command code
// if the string contains space
func commandDecoder(command string) string {
	splitString := strings.Split(command, " ")
	if len(splitString) > 1 {
		return splitString[0]
	}
	return command
}

func GetDefeaultAction() []CommandReaderOption {
	return []CommandReaderOption{
		Place,
		Move,
		Left,
		Right,
		Report,
	}
}
