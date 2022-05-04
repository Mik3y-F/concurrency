package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 15; i++ {
		wg.Add(3)
		go func() {
			go func() {
				fmt.Println("task 1")
				wg.Done()
			}()
			go func() {
				fmt.Println("task 2")
				wg.Done()
			}()
			go func() {
				fmt.Println("task 3")
				wg.Done()
			}()
			wg.Wait()
		}()
	}
}
