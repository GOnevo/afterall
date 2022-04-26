package main

import (
	"github.com/gonevo/afterall"
	"syscall"
)

type Application struct{}

func (a *Application) Shutdown() {}

func main() {

	app := &Application{}
	// here is code of your awesome application

	f := func() {
		// you can log
		app.Shutdown()
		// or send notifications
	}
	afterall.I().HaveToCall(f).On(syscall.SIGTERM, syscall.SIGINT).Wait()
}
