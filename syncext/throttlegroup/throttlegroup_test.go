package throttlegroup_test

import (
	"fmt"
	"time"

	"github.com/andreynering/goext/syncext/throttlegroup"
)

func ExampleGroup() {
	g := throttlegroup.WithThrottle(2)

	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() error {
			fmt.Println(i)
			time.Sleep(time.Second)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		panic(err)
	}
}
