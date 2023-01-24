package main

import (
	"fmt"
	"math"
)

func main() {

	const a = 50000000000

	fmt.Println(a)

	const b = 3e40 / a

	fmt.Println(b)

	fmt.Println(math.Sin(b))
	fmt.Println(math.Sin(22 / 7.0))

}
