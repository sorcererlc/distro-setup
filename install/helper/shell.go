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

func Run(a ...string) error {
	b, a := a[0], a[1:]
	c := Cmd{
		Bin:  b,
		Args: append([]string{}, a...),
	}

	return executeCommand(c)
}

func RunStdin(a ...string) ([]byte, error) {
	b, a := a[0], a[1:]

	r, err := exec.Command(b, a...).Output()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func executeCommand(cmd Cmd) error {
	l := log.NewStdOutLog()

	if os.Getenv("TEST") == "true" {
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
