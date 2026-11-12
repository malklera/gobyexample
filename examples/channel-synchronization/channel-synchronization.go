// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,
// you may prefer to use a [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second*5)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {

	t := time.Now()
	fmt.Println("Begining:", t.Format(time.TimeOnly))

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)

	// This happens without waiting for `worker`
	fmt.Println("Before reveive from done:", time.Since(t))

	// Block until we receive a notification from the
	// worker on the channel.
	<-done

	fmt.Println("After reveive from done:", time.Since(t))
}
