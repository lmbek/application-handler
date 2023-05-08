package application

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

type ChromeLauncher struct {
	Cmd *exec.Cmd
}

// Start - starts the chrome launcher, graceful shutdowns should not be neccesary,
// as we only open or close the frontend here, no file writing or similar is happening
func (chromeLauncher *ChromeLauncher) Start() error {
	// Start frontend by starting a new Chrome process
	path := os.Getenv("programfiles") + "\\Google\\Chrome\\Application\\chrome.exe"
	frontendInstallationPath := os.Getenv("localappdata") + "\\Google\\Chrome\\InstalledApps\\" + "DefaultOrganisationName" + "\\" + "DefaultProjectName"
	chromeLauncher.Cmd = exec.Command(path, "--app=https://google.dk", "--user-data-dir="+frontendInstallationPath)
	err := chromeLauncher.Cmd.Start()
	if err != nil {
		println("Warning: Chrome could not start, is it installed?")
		return err
	}

	signalHandler := make(chan os.Signal, 1)
	signal.Notify(signalHandler, os.Interrupt)
	go func() {
		<-signalHandler // waiting for termination
		err := chromeLauncher.Shutdown()
		if err != nil {
			fmt.Println(err)
		}

	}()

	err = chromeLauncher.Cmd.Wait() // waiting for exiting program
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Shutdown - Shuts down the chromeLauncher, if it has started, otherwise gives "chrome process is not running"
func (chromeLauncher *ChromeLauncher) Shutdown() error {
	if chromeLauncher.Cmd == nil || chromeLauncher.Cmd.Process == nil {
		return errors.New("chrome process is not running")
	}
	err := chromeLauncher.Cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}
