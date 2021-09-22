package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1024)

	// imitate timeout method, preveting program from waiting too long
	timeout := time.After(time.Second)

	go func() {
		time.Sleep(5 * time.Millisecond)
		ch <- 100
	}()

	select {
	case <-ch:
		fmt.Println("ch")
	case <-timeout:
		fmt.Println("timeout")
	}
}
