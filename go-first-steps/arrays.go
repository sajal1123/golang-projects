package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] += i
	}
	fmt.Println("arr a ->", arr)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr b ->", b)

	var mat [2][3]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			mat[i][j] += (i + j)
		}
	}
	fmt.Println("mat ->", mat)
}
