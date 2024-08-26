package helper

import (
	"os"
	"os/exec"
	"strings"
)

type Cmd struct {
	Bin  string
	Args []string
}

func ExecuteCommand(cmd Cmd) error {
	println("Executing '" + cmd.Bin + " " + strings.Join(cmd.Args, " ") + "'")
	return nil
	c := exec.Command(cmd.Bin, cmd.Args...)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
