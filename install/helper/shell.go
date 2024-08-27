package helper

import (
	"os"
	"os/exec"
	"setup/log"
	"strings"
)

type Cmd struct {
	Bin  string
	Args []string
}

func ExecuteCommand(cmd Cmd) error {
	l := log.NewStdOutLog()

	if os.Getenv("DRY_RUN") == "true" {
		l.Debug("Executing " + log.Cyan + cmd.Bin + " " + strings.Join(cmd.Args, " ") + log.Reset)
		return nil
	}

	c := exec.Command(cmd.Bin, cmd.Args...)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
