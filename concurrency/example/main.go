package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")

	jobs := 5

	quit := make(chan interface{})
    
    // fan-in chan pattern
	terminated := make(chan int, jobs)
    
    // run multiple jobs concurrently
	for i := 0; i < jobs; i++ {
	    // receive termination signal from each concurrent procs
		go func(i int) {
		    terminated <- <-proc(quit, i)
		}(i)
	}
    
    // a goroutine that simulates quit after some time
    // could actually be an error, force quit, etc.
	go func() {
		time.Sleep(5 * time.Second)
		close(quit)
	}()

	for i := 0; i < jobs; i++ {
		fmt.Printf("proc:%d done\n", <-terminated)
	}

	close(terminated)
}

// generator pattern
func proc(quit <-chan interface{}, id int) (<-chan int) {
	
	terminated := make(chan int) 
	
	go func() {

		defer func() {
			fmt.Printf("proc:%d exit signal received. running cleanup code...\n", id)
			terminated <- id
		}()

		for {
			select {
			case <-quit:
				return
			case <-time.After(time.Second):
				fmt.Printf("running proc:%d\n", id)
			}
		}
	}()
	
	return terminated
}