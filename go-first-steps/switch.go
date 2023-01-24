package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()

	switch {
	case a.Hour() < 12:
		fmt.Println("it is before noon")
	case a.Hour() >= 12:
		fmt.Println("it is after noon")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("WEEKEND")
	default:
		fmt.Println("weekday")
	}
	typeIdentifier := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("BOOLEAN")
		case string:
			fmt.Println("STRING")
		case int:
			fmt.Println("INTEGER")
		default:
			fmt.Println("TYPE UNDEFINED", i)
		}
	}
	typeIdentifier("Hell")
	typeIdentifier(12)
	typeIdentifier(12.1)
	typeIdentifier(true)
}
