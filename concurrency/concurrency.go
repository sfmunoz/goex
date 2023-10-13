//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/goex
// Date: Fri Oct 13 05:06:59 AM UTC 2023
//

// {{{ package

package concurrency

// }}}
// {{{ imports

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// }}}
// ---- funcs (private) ----
// {{{ func slumber()

func slumber(ctx context.Context, wg *sync.WaitGroup, t time.Duration) {
	fmt.Println("==== slumber() ====")
	defer fmt.Println("---- slumber() ----")
	defer (*wg).Done()
	select {
	case <-time.After(t):
		fmt.Println("slumber(): nap finished after", t)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("slumber(): err =", err, "-->", ctx)
	}
}

// }}}
// {{{ func push()

func push(c *chan int, x int) {
	time.Sleep(500 * time.Millisecond)
	*c <- x
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	fmt.Println("==== begin ====")
	// waitgroup, context, ...
	wg := sync.WaitGroup{}
	delays := []time.Duration{500 * time.Millisecond, 1500 * time.Millisecond}
	for _, delay := range delays {
		// alt: 'ctx := context.TODO()'
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
		defer cancel() // use 'go vet' to make sure 'cancel()' is properly used
		wg.Add(1)
		go slumber(ctx, &wg, delay)
		wg.Wait()
	}
	// goroutines, channels, ...
	c := make(chan int) // unbuffered; buffered: "make(chan int, 10)"
	go func() {
		wg.Add(1)
		defer wg.Done() // == wg.Add(-1)
		c <- 111
		time.Sleep(1200 * time.Millisecond)
	}()
	go push(&c, 222)
	for i := 0; i < 2; i++ {
		n := <-c
		fmt.Printf("[%d] n = %d\n", i, n)
	}
	wg.Wait()
	fmt.Println("---- end ----")
}

// }}}
