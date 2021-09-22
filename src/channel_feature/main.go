package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	id int
	in chan int
	wg *sync.WaitGroup
}

func work(worker Worker) {
	for data := range worker.in {
		fmt.Printf("worker %d received %c\n", worker.id, data)
		worker.wg.Done()
	}
}

func createWork(id int) {
	worker := Worker{
		id: 0,
		in: make(chan int),
		wg: &sync.WaitGroup{},
	}

	go work(worker)

	for i := 0; i < 10; i++ {
		worker.wg.Add(1)
		worker.in <- 'a' + i
	}

	worker.wg.Wait()
}

func main() {
	createWork(0)
}
