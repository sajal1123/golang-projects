package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "सजल"
	fmt.Println("len = ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d character = %x\n", i+1, s[i])
	}
	fmt.Println("Rune count = ", utf8.RuneCountInString(s))

	for idx, rval := range s {
		fmt.Printf("%#U starts at %d\n", rval, idx)
	}
}
