package main

import "fmt"

func main() {
	nums := [3]int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of nums =", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Num 3 is at index", i)
		}
	}

	m := map[string]string{"a": "Awesome", "b": "Brilliant"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "sajal" {
		fmt.Println(i, c)
	}
}
