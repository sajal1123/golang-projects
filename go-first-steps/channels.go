package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping1"
		messages <- "ping2"
		messages <- "ping3"
	}()

	//msg := <-messages
	//fmt.Println(msg)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
