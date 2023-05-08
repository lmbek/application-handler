package tests

import (
	"application"
	"errors"
	"fmt"
	"helpers"
	"sync"
	"testing"
	"time"
)

func TestChromeLauncher_Start(test *testing.T) {
	test.Run("Testing ChromeLauncher Start", func(test *testing.T) {
		app := new(application.ChromeLauncher)
		waitgroup := &sync.WaitGroup{}

		go func() {
			// shutdown app after 3 seconds
			time.Sleep(time.Second * 3)
			err := app.Shutdown()
			if err != nil {
				fmt.Println(err)
			}

			waitgroup.Done()
		}()

		waitgroup.Add(1)
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
		done := make(chan struct{})

		go func() {
			// shutdown app after 3 seconds
			time.Sleep(time.Second * 3)
			result := app.Shutdown()
			if result != nil {
				fmt.Println(result)
			}
			close(done)
		}()

		waitgroup.Add(1)
		go func() {
			defer waitgroup.Done()
			<-done
		}()

		result := app.Start()
		expected := error(nil)
		helpers.StandardTestChecking(test, result, expected)
		waitgroup.Wait()
	})

	test.Run("Testing ChromeLauncher Shutdown Without Starting It First", func(test *testing.T) {
		app := new(application.ChromeLauncher)
		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()
		}()

		result := app.Shutdown()
		expected := errors.New("chrome process is not running")
		helpers.StandardTestChecking(test, result.Error(), expected.Error())
		wg.Wait()
	})
}
