package main

import (
	"fmt"
	"sync"
	"time"
)


var m=sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Println("The final time is", time.Since(t0))
	fmt.Println("The final result table is", results)

}

func dbCall(i int) {
	var delay time.Duration = 2000 * time.Millisecond
	time.Sleep(delay)
	fmt.Printf("the result from the database is %v\n", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()

	wg.Done()
}
