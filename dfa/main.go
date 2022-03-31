package main

import (
	"fmt"
	"time"
)

type Data struct {
	Values []float64
}

type Component struct {
	ID string

	To *Component

	DataChan chan *Data

	Quit <-chan interface{}
	Done chan string
}

type IncrementComponent struct {
	Component
}

func (c *Component) Cleanup() {
	fmt.Printf("component[%s] running cleanup code...\n", c.ID)
}

func (c *IncrementComponent) Perform(data *Data) {
	newData := &Data{}

	for i := 0; i < len(data.Values); i++ {
		newData.Values = append(newData.Values, data.Values[i]+1)
	}

	c.To.DataChan <- newData
}

func (c *IncrementComponent) Run() {
	c.Done = make(chan string)

	go func() {
		defer func() {
			c.Cleanup()
			c.Done <- c.ID
		}()

		for {
			select {
			case <-c.Quit:
				return
			case data := <-c.DataChan:
				c.Perform(data)
			}
		}
	}()
}

type PrinterComponent struct {
	Component
}

func (c *PrinterComponent) Run() {
	c.Done = make(chan string)

	go func() {
		defer func() {
			c.Cleanup()
			c.Done <- c.ID
		}()

		for {
			select {
			case <-c.Quit:
				return
			case data := <-c.DataChan:
				c.Perform(data)
			}
		}
	}()
}

func (c *PrinterComponent) Perform(data *Data) {
	fmt.Printf("component[%s]: %v\n", c.ID, data.Values)
}

func main() {
	fmt.Println("Data flow apllication example")

	quit := make(chan interface{})
	done := make(chan string, 3)

	c1 := &IncrementComponent{
		Component: Component{
			ID:       "C1",
			DataChan: make(chan *Data),
			Quit:     quit,
		},
	}

	c2 := &IncrementComponent{
		Component: Component{
			ID:       "C2",
			DataChan: make(chan *Data),
			Quit:     quit,
		},
	}

	c3 := &PrinterComponent{
		Component: Component{
			ID:       "C3",
			DataChan: make(chan *Data),
			Quit:     quit,
		},
	}

	c1.To = &c2.Component
	c2.To = &c3.Component

	c1.Run()
	c2.Run()
	c3.Run()

	go func() {
		done <- <-c1.Done
	}()

	go func() {
		done <- <-c2.Done
	}()

	go func() {
		done <- <-c3.Done
	}()

	go func() {
		<-time.After(5 * time.Second)
		close(quit)
	}()

	go func() {
		<-time.After(2 * time.Second)
		c1.DataChan <- &Data{[]float64{1, 2, 3, 4, 5}}
	}()

	for i := 0; i < 3; i++ {
		fmt.Printf("component[%s] done\n", <-done)
	}

	close(done)
}
