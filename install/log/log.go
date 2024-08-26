package log

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Terminal color
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func ts() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func args(args ...string) string {
	s := ""
	for _, a := range args {
		s += " " + a
	}

	return s
}

func writeLogStdout(t string, ll string, s string, c string) {
	fmt.Printf("[%s] %s%s%s %s\n", t, c, ll, Reset, s)
}

func ClearLogs() {
	c := exec.Command("rm", "-rf", "logs")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin

	err := c.Run()
	if err != nil {
		writeLogStdout(ts(), "ERROR", err.Error(), Red)
		os.Exit(1)
	}
}

type Log struct {
	StdOutOnly bool
	LogFile    string
}

func NewStdOutLog() *Log {
	return &Log{
		StdOutOnly: true,
		LogFile:    "",
	}
}

func NewLog(f string) *Log {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		writeLogStdout(ts(), "ERROR", err.Error(), Red)
		os.Exit(1)
	}

	return &Log{
		StdOutOnly: false,
		LogFile:    "logs/" + f,
	}
}

func (l *Log) writeLogFile(t string, ll string, s string) {
	if l.StdOutOnly {
		return
	}
	fs, err := os.OpenFile(l.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		writeLogStdout(ts(), "ERROR", err.Error(), Red)
		os.Exit(1)
	}

	defer fs.Close()

	_, err = fs.WriteString(t + " " + ll + " " + s + "\n")
}

func (l *Log) Info(a ...string) {
	t := ts()
	ll := "INFO"
	s := args(a...)

	l.writeLogFile(t, ll, s)
	writeLogStdout(t, ll, s, Green)
}

func (l *Log) Error(a ...string) {
	t := ts()
	ll := "ERROR"
	s := args(a...)

	l.writeLogFile(t, ll, s)
	writeLogStdout(t, ll, s, Red)
}
