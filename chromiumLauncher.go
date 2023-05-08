package application

import (
	"errors"
	"os/exec"
)

type ChromiumLauncher struct {
	Cmd exec.Cmd
}

func (chromiumLauncher ChromiumLauncher) Start() error {
	//TODO implement
	return errors.New("not implemented")
}

func (chromiumLauncher ChromiumLauncher) Shutdown() error {
	//TODO implement
	return errors.New("not implemented")
}
