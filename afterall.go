package afterall

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// Afterall closes your application correctly
type Afterall struct {
	signals    []os.Signal
	fn         func()
	wg         *sync.WaitGroup
	sigChannel chan os.Signal
	nowChannel chan bool
	Error      error
}

// I returns new Afterall instance.
func I() *Afterall {
	i := &Afterall{
		sigChannel: make(chan os.Signal, 1),
		nowChannel: make(chan bool, 1),
		wg:         &sync.WaitGroup{},
	}

	return i
}

// HaveToCall accepts a function that will be called after all.
func (a *Afterall) HaveToCall(f func()) *Afterall {
	a.fn = f
	return a
}

// On allows you to set the necessary signals.
func (a *Afterall) On(signals ...os.Signal) *Afterall {
	a.signals = signals
	return a
}

// Now calls HaveToCall function immediately.
func (a *Afterall) Now() {
	a.nowChannel <- true
	a.Wait()
}

// Wait for signal and do what I have to do.
func (a *Afterall) Wait() {
	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			a.Error = fmt.Errorf("%s", recoveryMessage)
		}
	}()
	a.wg.Add(1)
	a.attachSignals()
	go a.listenForSignal()
	a.wg.Wait()
	if a.fn != nil {
		a.fn()
	}
}

func (a *Afterall) attachSignals() {
	signal.Notify(a.sigChannel, a.signals...)
}

func (a *Afterall) listenForSignal() {
	defer a.wg.Done()
	for {
		select {
		case <-a.sigChannel:
			return
		case <-a.nowChannel:
			return
		}
	}
}
