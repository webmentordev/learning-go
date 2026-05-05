package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{} // Mutex become RWMutex
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

/*
RWMutex has additional RLock() and RUnlock() methods
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
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock() // This sets full lock, all RLock and Locks must be cleared
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() /* Aquire the read lock if there is no full lock.
	It will not read unless the full lock exist or mean data ins being added
	in the slice */
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
