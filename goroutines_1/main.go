package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

/*
add 'go' before the function to run the function concurrently
but we need to create 'waitGroup' which is like a counter.
It increments when a function is running and want until counter
is 0 to complete the program execution. 0 waitGroup means
that all tasks are completed.

what if we want to pass the data to a slice. see goroutines_2 folder
*/
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // Increment waitGroup before spawning goroutine
		go dbCall(i) // go here
	}
	wg.Wait() // Wait for the waitGroup to be 0
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	wg.Done() // We call Done method to decrement the waitgroup counter
}
