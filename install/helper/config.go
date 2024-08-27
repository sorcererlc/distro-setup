package helper

import (
	"os"
	"setup/log"
	"setup/types"

	"gopkg.in/yaml.v3"
)

func GetConfig(e *types.Environment) (*types.Config, error) {
	l := log.NewLog("config.log")
	conf := types.Config{}

	fs := e.Cwd + "/config.yml"
	l.Info("Loading config file", fs)

	f, err := os.ReadFile(fs)
	if err != nil {
		l.Error("Error loading config file", err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		l.Error("Error decoding config file", err.Error())
		return nil, err
	}

	return &conf, nil
}
