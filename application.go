package application

import (
	"errors"
	"os/exec"
)

// interfaces and creation of new instances that implement the interfaces

type Application struct {
	Cmd exec.Cmd
}

func (application *Application) Start() error {
	//TODO implement
	return errors.New("not implemented")
}

func (application *Application) Shutdown() error {
	//TODO implement
	return errors.New("not implemented")
}
