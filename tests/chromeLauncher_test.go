package tests

import (
	"application"
	"errors"
	"fmt"
	"helpers"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestChromeLauncher_Installed(test *testing.T) {
	test.Run("Testing ChromeLauncher Start", func(test *testing.T) {
		app := new(application.ChromeLauncher)
		waitgroup := &sync.WaitGroup{}
		waitgroup.Add(1)
		go func() {
			// shutdown app after 3 seconds
			time.Sleep(time.Second * 3)
			err := app.Shutdown()
			if err != nil {
				fmt.Println(err)
			}

			waitgroup.Done()
		}()

		result := app.Start()
		if result != nil {
			if strings.Contains(result.Error(), "file does not exist") && strings.Contains(result.Error(), "exec") {
				helpers.PrintGotFatalError(test, "Could not start chrome, is it installed?")
			}
		}

		expected := error(nil)
		helpers.StandardTestChecking(test, result, expected)
		waitgroup.Wait()
	})
}

func TestChromeLauncher_Start(test *testing.T) {
	test.Run("Testing ChromeLauncher Start", func(test *testing.T) {
		app := new(application.ChromeLauncher)
		waitgroup := &sync.WaitGroup{}
		waitgroup.Add(1)
		go func() {
			// shutdown app after 3 seconds
			time.Sleep(time.Second * 3)
			err := app.Shutdown()
			if err != nil {
				fmt.Println(err)
			}

			waitgroup.Done()
		}()

		result := app.Start()
		expected := error(nil)
		helpers.StandardTestChecking(test, result, expected)
		waitgroup.Wait()
	})
}

func TestChromeLauncher_Shutdown(test *testing.T) {
	test.Run("Testing ChromeLauncher Shutdown", func(test *testing.T) {
		app := new(application.ChromeLauncher)
		waitgroup := &sync.WaitGroup{}
		waitgroup.Add(1)

		go func() {
			// shutdown app after 3 seconds
			time.Sleep(time.Second * 3)
			result := app.Shutdown()
			if result != nil {
				fmt.Println(result)
			}
			waitgroup.Done()
		}()

		result := app.Start()
		expected := error(nil)
		helpers.StandardTestChecking(test, result, expected)
		waitgroup.Wait()
	})

	test.Run("Testing ChromeLauncher Shutdown Without Starting It First", func(test *testing.T) {
		app := new(application.ChromeLauncher)
		waitgroup := &sync.WaitGroup{}

		waitgroup.Add(1)
		go func() {
			time.Sleep(time.Second * 2)
			defer waitgroup.Done()
		}()

		result := app.Shutdown()
		expected := errors.New("chrome process is not running")
		helpers.StandardTestChecking(test, result.Error(), expected.Error())
		waitgroup.Wait()
	})
}
