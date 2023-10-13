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
// {{{ func count()

func count(s string, c chan<- string) {
    // chan ..... send/receive
    // chan<- ... send-only
    // <-chan ... receive-only
	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(500 * time.Millisecond)
		}
		c <- fmt.Sprintf("%s-%d", s, i)
	}
	close(c)
}

// }}}
// {{{ func main1()

func main1() {
	fmt.Println("==== main1() ====")
	defer fmt.Println("---- main1() ----")
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
}

// }}}
// {{{ func main2()

func main2() {
	fmt.Println("==== main2() ====")
	defer fmt.Println("---- main2() ----")
	wg := sync.WaitGroup{}
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
}

// }}}
// {{{ func main3()

func main3() {
	fmt.Println("==== main3() ====")
	defer fmt.Println("---- main3() ----")
	c := make(chan string)
	go count("item", c)
	for x := range c {
		fmt.Println(x)
	}
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	fmt.Println("==== Main() ====")
	defer fmt.Println("---- Main() ----")
	main1()
	main2()
	main3()
}

// }}}
