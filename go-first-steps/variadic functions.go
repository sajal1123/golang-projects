package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
	tot := 0
	for _, num := range nums {
		tot += num
	}
	fmt.Println(tot)
}

func main() {
	sum(1, 2, 3)

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 3e10)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum(nums...)
}