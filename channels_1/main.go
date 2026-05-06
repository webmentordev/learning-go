package main

import (
	"fmt"
	"time"
)

/*
- Channels hold data, data can a slice, int e.t.c.
- Thread safety as they avoid data races during read/write
- Listen to when data is added/removed then block code execution

Channels are properly used with goroutines
*/

func main() {
	/*
		var c = make(chan int) // Declare a unbuffered channel (single value)
		c <- 1                 // passing data to a channel
		var i = <-c            // moved the value from c and assing to i
		fmt.Println(i)
		// This code will casue deadlock as c <- 1 and nothing read from it
	*/

	var c = make(chan int)
	go process(c)
	fmt.Println(<-c) // Reading the value

	var a = make(chan int)
	go loop_process(a)
	for i := range a {
		// loop waits until value is added to the channel
		// It will casue deadlock when waiting for 5th value
		fmt.Println(i)
	}

	var b = make(chan int, 5) // this is buffered channel
	go buffered_process(b)
	// buffered_process will finish instantly but loop will keep running and reading the vales in the channel
	for z := range b {
		fmt.Println(z)
		time.Sleep(time.Second * 1)
	}

}

func process(c chan int) {
	c <- 123
}

func loop_process(a chan int) {
	defer close(a) // It means do this before function exists, or use close
	for i := 0; i < 5; i++ {
		a <- i
	}
	// close(a) Close so loop does not casue deadlock with 5th value
}

func buffered_process(b chan int) {
	defer close(b)
	for i := 0; i < 5; i++ {
		b <- 1
	}
	fmt.Println("Exiting process")
}
