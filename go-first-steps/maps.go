package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["x"] = 0
	m["y"] = 1
	fmt.Println("Map -> ", m)

	fmt.Println("First k-v pair ->", m["x"])

	delete(m, "x")
	fmt.Println("Map after deletion -> ", m)

	nm := map[string]int{"a": 1, "b": 12}
	fmt.Println("New Map ->", nm)
}
