package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done() // this way done is always called no matter what happens in the go routine
			work(i + 1)
		}(i)

	}

	wg.Wait()

	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task:", id, " - is done")
}
