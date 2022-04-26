# Introduction

Afterall is a library with a user-friendly interface for correct processing of your application closing.

## Installation

```shell
go get github.com/gonevo/afterall
```

## Usage
All you need to do is write one line and make sure it's at the end of your main function.

## Example

```go
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
		// you can log ...
		app.Shutdown()
		// ... or send notifications
	}
	
	afterall.I().HaveToCall(f).On(syscall.SIGTERM, syscall.SIGINT).Wait()
}
```

## Methods

##### `Afterall.I()` returns new configurable instance of Afterall
##### `Afterall.On()` allows you to set the necessary signals
##### `Afterall.HaveToCall()` accepts a function that will be called after all
##### `Afterall.Wait()` starts listening to signals
##### `Afterall.Now()` calls HaveToCall function immediately

## License

The afterall library is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
