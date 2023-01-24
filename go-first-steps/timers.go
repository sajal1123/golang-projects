package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 fired.")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired.")
	}()
	stopper := timer2.Stop()
	if stopper {
		fmt.Println("Timer 2 stopped.")
	}
	time.Sleep(2 * time.Second)
}
