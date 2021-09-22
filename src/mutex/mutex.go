package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	val  int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.val++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.val)
}

func main() {
	a := &atomicInt{val: 0}
	a.increment()
	go a.increment()
	time.Sleep(1 * time.Second)
	fmt.Println(a.get())
}
