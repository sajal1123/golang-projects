package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("slice->", s)

	s = append(s, "d")
	s = append(s, "e", "f", string('G'))
	fmt.Println("Appended slice->", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy slice->", c)

	l := s[2:]
	fmt.Println("s[2:] = ", l)
	fmt.Println("s[:3] = ", s[:3])
	fmt.Println("s[1:3] = ", s[1:3])

	mat := make([][]int, 3)
	for i := 0; i < 3; i++ {
		x := i + 1
		mat[i] = make([]int, x)
		for j := 0; j < x; j++ {
			mat[i][j] = i + j
		}
	}
	fmt.Println("Mat -> ", mat)

}
