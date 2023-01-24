package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("10! = ", fact(10))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("Fib(11) = ", fib(11))
}
