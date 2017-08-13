# Throttle

[![GoDoc](https://godoc.org/github.com/andreynering/goext/syncext/throttle?status.svg)](https://godoc.org/github.com/andreynering/goext/syncext/throttle)

Package throttle is a simple utility that helps you throttle the number
of running goroutines when using [sync.WaitGroup][waitgroup].
This is useful while running CPU and/or memory intensive code concurrently
(e.g.: if you launch too many goroutines, you will end up of resources).

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/andreynering/goext/syncext/throttle"
)

func main() {
	var wg sync.WaitGroup
	th := throttle.New(2)

	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer th.Done()
			th.Wait()

			fmt.Println(i)
			time.Sleep(time.Second)
		}()
	}

	wg.Wait()
}
```

[waitgroup]: https://golang.org/pkg/sync/#WaitGroup
