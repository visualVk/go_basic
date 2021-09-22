package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for ch := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, ch)
	}
}

func createWorker() chan<- int {
	c := make(chan int)
	go worker(0, c)
	return c
}

// c1 and c2 generator data, called source of message
// worker is used for receive message from c1 and c2
// but we also need cache of message preventing message from replacing
// so creating values to cache message
func main() {
	c1, c2 := generator(), generator()
	worker := createWorker()

	var values []int
	tm := time.After(8 * time.Second)
	tk := time.Tick(time.Second)

	for {
		// create nil channel that can recevie message, but cannnot handle message
		var activeChannel chan<- int
		var activeValue int
		// when len(value) > 0 that means some message had been stored in values,
		// we should send data from values to worker for consuming message
		if len(values) > 0 {
			activeChannel = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeChannel <- activeValue:
			values = values[1:]
		// handle timeout
		case <-time.After(8 * time.Second):
			fmt.Println("timeout")
		// check len(values) per second
		case <-tk:
			fmt.Printf("length of values: %d\n", len(values))
		// exit when arriving to 8s
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
