package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "value passed after 2 seconds channel 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res, "has been received.")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout channel 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "value passed after 2 seconds channel 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res, "has been received.")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout channel 2")
	}
}
