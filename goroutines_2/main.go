package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{} // Add Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

/*
If we want to pass the data to a slice, we can Mutax to lock the goroutine,
modify then pass the data. If other goroutines has a lock then current
goroutine will wait for release.
It matters where you put your local() & unlock(), it can effect the concurrency

Mutex locks out other goroutines from accessing the result slice, see goroutines_3 folder for more RWMutex
*/
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock() // Lock the goroutine if not held by other goroutine
	results = append(results, dbData[i])
	m.Unlock() // Release the lock after processing so other goroutines can process
	wg.Done()
}
