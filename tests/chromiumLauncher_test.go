package tests

import (
	"application"
	"errors"
	"helpers"
	"testing"
)

func TestChromiumLauncher_Start(test *testing.T) {
	test.Run("Testing ChromiumLauncher Start", func(test *testing.T) {
		app := new(application.ChromiumLauncher)
		result := app.Start().Error()
		expected := errors.New("not implemented").Error()
		helpers.StandardTestChecking(test, result, expected)
	})
}

func TestChromiumLauncher_Shutdown(test *testing.T) {
	test.Run("Testing ChromiumLauncher Shutdown", func(test *testing.T) {
		app := new(application.ChromiumLauncher)
		result := app.Start().Error()
		expected := errors.New("not implemented").Error()
		helpers.StandardTestChecking(test, result, expected)
	})
}
