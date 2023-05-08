package application

type LauncherInterface interface {
	Start() error
	Shutdown() error
}
