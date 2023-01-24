package main

import "fmt"

func insSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := insSeq()

	for i := 0; i < 11; i++ {
		fmt.Println(nextInt())
	}

	fmt.Println(nextInt())
	nextInt2 := insSeq()
	fmt.Println(nextInt2())
}
