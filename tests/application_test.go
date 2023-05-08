package tests

import (
	"application"
	"errors"
	"helpers"
	"testing"
)

func TestApplication_Start(test *testing.T) {
	// maybe we will remove this - as start is not exported
	test.Run("Testing Application Start", func(test *testing.T) {
		app := new(application.Application)
		result := app.Start().Error()
		expected := errors.New("not implemented").Error()
		helpers.StandardTestChecking(test, result, expected)
	})
}

func TestApplication_Shutdown(test *testing.T) {
	test.Run("Testing Application Shutdown", func(test *testing.T) {
		app := new(application.Application)
		result := app.Shutdown().Error()
		expected := errors.New("not implemented").Error()
		helpers.StandardTestChecking(test, result, expected)
	})
}
