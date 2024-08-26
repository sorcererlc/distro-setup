package helper

import (
	"os"
	"strings"
)

type OS struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"pretty_name"`
	Version    string `json:"version"`
	VersionId  string `json:"version_id"`
}

func parseRelease() (map[string]string, error) {
	o := make(map[string]string)

	r, err := os.ReadFile("/etc/os-release")
	if err != nil {
		panic(err)
	}

	l := strings.Split(string(r), "\n")

	for _, i := range l {
		if i == "" {
			break
		}
		v := strings.Split(i, "=")
		o[v[0]] = strings.ReplaceAll(v[1], "\"", "")
	}

	return o, nil
}

func GetOS() (*OS, error) {
	r, err := parseRelease()
	if err != nil {
		return nil, err
	}

	o := OS{
		Id:         r["ID"],
		Name:       r["NAME"],
		PrettyName: r["PRETTY_NAME"],
		Version:    r["VERSION"],
		VersionId:  r["VERSION_ID"],
	}

	return &o, nil
}
