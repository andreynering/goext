// Package throttle is a simple utility that helps you throttle the number
// of running goroutines when using sync.WaitGroup.
// This is useful while running CPU and/or memory intensive code concurrently
// (e.g.: if you launch too many goroutines, you will end up of resources).
package throttle

import (
	"runtime"
)

// Throttle is the struct used for throttling.
// The zero value meansno throttling.
type Throttle struct {
	ch chan struct{}
}

// Default return throttle sized by the number of CPU cores of the device.
func Default() Throttle {
	return New(runtime.NumCPU())
}

// New returns throttle sized by the given number.
// E.g.: throttle of 4 means 4 active goroutines at the same time.
func New(num int) Throttle {
	t := Throttle{
		ch: make(chan struct{}, num),
	}
	for i := 0; i < num; i++ {
		t.ch <- struct{}{}
	}
	return t
}

// Wait waits til the next goroutine is done to start its job.
func (t Throttle) Wait() {
	if t.ch != nil {
		<-t.ch
	}
}

// Done should be called when a goroutine finished it's job.
func (t Throttle) Done() {
	if t.ch != nil {
		t.ch <- struct{}{}
	}
}
