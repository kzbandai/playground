package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	startCh := make(chan int, 2)
	endCh := make(chan int)

	for i := 0; i < 2; i++ {
		go func(integer int) {
			t := rand.Intn(19)
			fmt.Println("rand number: ", t)
			time.Sleep(time.Duration(t) * time.Second)
			startCh <- t
		}(i)
	}

	go func() {
		select {
		case sig := <-startCh:
			fmt.Println("receive int: ", sig)
			endCh <- 1
		}
	}()
	fmt.Println("fugafuga")

	<-endCh

	println("receive int: ", <-startCh)

	//println("receive int: ", <-startCh) // panic
}
