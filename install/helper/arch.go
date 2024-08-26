package helper

import (
	"errors"
	"fmt"
	"strings"
)

// TODO Make this whole thing work

type ArchHelper struct{}

func (f *ArchHelper) checkInstalledPackage(p string) bool {
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

func (f *ArchHelper) UpdateDistro() error {
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

func (f *ArchHelper) InstallPackages(p []string) error {
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

func (f *ArchHelper) RemovePackages(p []string) error {
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
