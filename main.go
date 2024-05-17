package main

import (
	"bufio"
	"os"

	"github.com/psbernardo/toyrobot/reader"
)

func main() {

	reader := reader.NewCommandReader(
		reader.Place,
		reader.Move,
		reader.Right,
		reader.Left,
		reader.Report,
		reader.Help,
		// add another command handler here
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		reader.RunCommand(scanner.Text())
	}

}
