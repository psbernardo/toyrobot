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
	TableTop           *pkg.TableTop
	mapCommandExecuter map[string]ActionHandler
}

func NewCommandReader(options ...CommandReaderOption) *CommandReader {
	c := &CommandReader{
		mapCommandExecuter: make(map[string]ActionHandler),
		TableTop:           pkg.NewTableTop(4, 4),
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
