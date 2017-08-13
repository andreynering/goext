# Throttle Group

[![GoDoc](https://godoc.org/github.com/andreynering/goext/syncext/throttlegroup?status.svg)](https://godoc.org/github.com/andreynering/goext/syncext/throttlegroup)

Package throttlegroup is a variation of the errgroup package
([golang/x/sync/errgroup][errgroup]) that uses throttling to control the active
number of goroutines at the same time.
This is useful while running CPU and/or memory intensive code concurrently
(e.g.: if you launch too many goroutines, you will end up of resources).

```go
package main

import (
	"fmt"
	"time"

	"github.com/andreynering/goext/syncext/throttlegroup"
)

func main() {
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
```

[errgroup]: https://godoc.org/golang.org/x/sync/errgroup
