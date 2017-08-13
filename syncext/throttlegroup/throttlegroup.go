// Package throttlegroup is a variation of the errgroup package
// (golang/x/sync/errgroup) that uses throttling to control the active
// number of goroutines at the same time.
// This is useful while running CPU and/or memory intensive code concurrently
// (e.g.: if you launch too many goroutines, you will end up of resources).
package throttlegroup

import (
	"context"
	"runtime"

	"github.com/andreynering/goext/syncext/throttle"

	"golang.org/x/sync/errgroup"
)

// Group is a throttle group. The zero value should not be used.
type Group struct {
	*errgroup.Group
	throttle throttle.Throttle
}

// Default returns a throttle group sized by the number of CPU cores of
// the device.
func Default() *Group {
	return WithThrottle(runtime.NumCPU())
}

// New returns a throttle group sized by the given number.
// E.g.: throttle of 4 means 4 active goroutines at the same time.
func WithThrottle(num int) *Group {
	return &Group{
		Group:    &errgroup.Group{},
		throttle: throttle.New(num),
	}
}

// DefaultWithContext returns a throttle group and a context, given a parent
// context.
// The size of the group will be the number of the CPU cores of the device.
func DefaultWithContext(ctx context.Context) (*Group, context.Context) {
	return WithContext(ctx, runtime.NumCPU())
}

// WithContext returns a throttle group and a context, given a parent context
// and a throttle size.
func WithContext(ctx context.Context, num int) (g *Group, resultCtx context.Context) {
	g = &Group{
		throttle: throttle.New(num),
	}
	g.Group, resultCtx = errgroup.WithContext(ctx)
	return
}

// Go overrides the parent errgroup.Group Go method, to make the throttling
// automatic.
func (g *Group) Go(f func() error) {
	g.Group.Go(func() error {
		g.throttle.Wait()
		defer g.throttle.Done()

		return f()
	})
}
