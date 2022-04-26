//go:build linux || bsd || darwin

package afterall

import (
	"os"
	"syscall"
	"testing"
	"time"
)

func TestAfterall_HaveToDo(t *testing.T) {
	t.Run("Default function", func(t *testing.T) {
		a := I().On(syscall.SIGTERM)
		a.Now()
	})

	t.Run("Function with panic", func(t *testing.T) {
		a := I().HaveToCall(func() {
			panic("panic message")
		}).On(syscall.SIGTERM)
		a.Now()
		if a.Error.Error() != "panic message" {
			t.Fatal("panic is not covered")
		}
	})
}

func TestAfterall_On(t *testing.T) {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	t.Run("SIGTERM", func(t *testing.T) {
		done := false
		a := I().HaveToCall(func() {
			done = true
		}).On(syscall.SIGTERM)
		AssertNotDone(t, done)
		go a.Wait()
		time.Sleep(1 * time.Second)
		_ = proc.Signal(syscall.SIGTERM)
		time.Sleep(1 * time.Second)
		AssertDone(t, done)
	})

	t.Run("SIGINT", func(t *testing.T) {
		done := false
		a := I().HaveToCall(func() {
			done = true
		}).On(syscall.SIGINT)
		AssertNotDone(t, done)
		go a.Wait()
		time.Sleep(1 * time.Second)
		_ = proc.Signal(syscall.SIGINT)
		time.Sleep(1 * time.Second)
		AssertDone(t, done)
	})

	t.Run("Two signals", func(t *testing.T) {
		done := false
		a := I().HaveToCall(func() {
			done = true
		}).On(syscall.SIGINT, syscall.SIGTERM)
		AssertNotDone(t, done)
		go a.Wait()
		time.Sleep(1 * time.Second)
		_ = proc.Signal(syscall.SIGINT)
		time.Sleep(1 * time.Second)
		AssertDone(t, done)
	})

	t.Run("Any signal", func(t *testing.T) {
		done := false
		a := I().HaveToCall(func() {
			done = true
		})
		AssertNotDone(t, done)
		go a.Wait()
		time.Sleep(1 * time.Second)
		_ = proc.Signal(syscall.SIGINT)
		time.Sleep(1 * time.Second)
		AssertDone(t, done)
	})
}

func TestAfterall_Now(t *testing.T) {
	t.Run("Now", func(t *testing.T) {
		done := false
		a := I().HaveToCall(func() {
			done = true
		}).On(syscall.SIGTERM)
		AssertNotDone(t, done)
		a.Now()
		AssertDone(t, done)
	})
}

func AssertNotDone(t *testing.T, d bool) {
	if d != false {
		t.Error("Expected Afterall is not done")
	}
}

func AssertDone(t *testing.T, d bool) {
	if d != true {
		t.Error("Expected Afterall is done")
	}
}
