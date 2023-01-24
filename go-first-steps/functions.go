package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("plus(1,2) = ", res)

	ress := plusPlus(1, 2, 3)
	fmt.Println("plusPlus(1,2,3) = ", ress)

	x, y := multiRet(2, 3)
	fmt.Println("Multi return = ", x, y)
}

func plus(a, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return plus(a, b) + c
}
func multiRet(a, b int) (int, int) {
	return a, b
}
