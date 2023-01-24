package main

import (
	"fmt"
)

func sender(pings chan<- string, msg string) {
	fmt.Println("loading message in sender...")
	pings <- msg
}
func handover(sender <-chan string, receiver chan<- string) {
	msg := <-sender
	fmt.Println("transferring message to receiver...")
	receiver <- msg
}

func main() {
	s := make(chan string, 1)
	r := make(chan string, 1)

	sender(s, "Hello, World!")
	handover(s, r)
	fmt.Println(<-r)
}
