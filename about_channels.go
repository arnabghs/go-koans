package go_koans

import "fmt"

func aboutChannels() {
	ch := make(chan string, 2)

	assert(len(ch) == 0) // channels are like buffers

	ch <- "foo" // i mean, "metaphors are like similes"

	assert(len(ch) == 1) // they can be queried for queued items

	assert(<-ch == "foo") // items can be popped out of them

	assert(len(ch) == 0) // and len() always reflects the "current" queue status

	// the 'go' keyword runs a function-call in a new "goroutine"
	// which executes "concurrently" with the calling "goroutine"
	go func() {
		takeOut := <-ch
		fmt.Println(takeOut)

		takeOut2 := <-ch
		fmt.Println(takeOut2)

		takeOut3 := <-ch
		fmt.Println(takeOut3)

		assert(len(ch) == 0)
	}()

	ch <- "bar"   // this send will succeed
	ch <- "quux"  // there's enough room for this send too
	ch <- "extra" // but the buffer only has two slots
}
