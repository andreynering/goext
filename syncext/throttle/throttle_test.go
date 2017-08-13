package throttle_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/andreynering/goext/syncext/throttle"
)

func ExempleThrottle() {
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
