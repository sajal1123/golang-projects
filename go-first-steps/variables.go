package main

import "fmt"

func main() {
	var a = "initialize"
	fmt.Println(a)
	print(a)

	var b, c int = 1, 7
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "fancy initialization"
	fmt.Println(f)

	var g float32
	fmt.Println("float32", g)

	var h float64
	println("float64", h)

	var i string
	fmt.Println("string -> ", i)
}
