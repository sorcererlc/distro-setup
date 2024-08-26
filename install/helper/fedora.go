package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type FedoraHelper struct {
	CorpRepos struct {
		Copr []string `yaml:"copr"`
	}
}

func NewFedoraHelper() (*FedoraHelper, error) {
	f := FedoraHelper{}

	return &f, nil
}

func (f *FedoraHelper) checkInstalledPackage(p string) bool {
	c := Cmd{
		Bin:  "dnf",
		Args: []string{"list", "installed", p},
	}

	err := ExecuteCommand(c)
	if err != nil {
		return false
	}

	return true
}

func (f *FedoraHelper) UpdateDistro() error {
	fmt.Printf("Updating packages\n")

	c := Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "upgrade", "-y"},
	}

	err := ExecuteCommand(c)
	if err != nil {
		fmt.Printf("Error updating packages. Aborting setup.\n%s", err.Error())
		return err
	}

	return nil
}

func (f *FedoraHelper) InstallPackages(p []string) error {
	fmt.Printf("Installing packages %s\n", strings.Join(p, ", "))

	c := Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "install", "-y"},
	}
	c.Args = append(c.Args, p...)

	err := ExecuteCommand(c)
	if err != nil {
		return err
	}

	for _, pk := range p {
		i := f.checkInstalledPackage(pk)
		if !i {
			return errors.New("Package " + pk + " failed to install. Aborting setup.")
		}
	}

	return nil
}

func (f *FedoraHelper) RemovePackages(p []string) error {
	fmt.Printf("Removing packages %s\n", strings.Join(p, ", "))

	c := Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "remove", "-y"},
	}
	c.Args = append(c.Args, p...)

	err := ExecuteCommand(c)
	if err != nil {
		return err
	}

	return nil
}

func (f *FedoraHelper) EnableCoprRepos() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fs, err := os.ReadFile(cwd + "/packages/fedora/repos.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fs, &f.CorpRepos)

	c := Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "copr", "enable", "-y"},
	}
	c.Args = append(c.Args, f.CorpRepos.Copr...)

	err = ExecuteCommand(c)
	if err != nil {
		return err
	}

	return nil
}

func (f *FedoraHelper) InstallRepos(v string, r []string) error {
	for i := 0; i < len(r); i++ {
		r[i] = fmt.Sprintf(r[i], v)
	}
	fmt.Printf("Installing repositories %s\n", strings.Join(r, ", "))

	c := Cmd{
		Bin:  "sudo",
		Args: []string{"dnf", "install", "-y"},
	}
	c.Args = append(c.Args, r...)

	err := ExecuteCommand(c)
	if err != nil {
		return err
	}

	return nil
}
