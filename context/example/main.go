package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Context in Go")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	jobs := 5
	procs := make(chan string, jobs)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	for i := 1; i <= jobs; i++ {
		go func(id int) {
			procs <- <-proc(ctx, "proc-"+strconv.Itoa(id))
		}(i)
	}

	select {
	case <-c:
		cancel()
	}

	for i := 0; i < jobs; i++ {
		fmt.Printf("main: completed [%s]\n", <-procs)
	}
}

func proc(ctx context.Context, id string) <-chan string {
	done := make(chan string)

	cleanup := func() {
		fmt.Printf("[%s] cleanup\n", id)
		time.Sleep(time.Duration(rand.Int63n(3000)) * time.Millisecond)
		done <- id
	}

	go func() {
		defer cleanup()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("\n[%s] received ctx cancellation\n", id)
				return
			default:
				fmt.Printf("[%s] running\n", id)
			}
			time.Sleep(time.Duration(rand.Int63n(250)) * time.Millisecond)
		}
	}()

	return done
}
