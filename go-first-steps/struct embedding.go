package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	fmt.Printf("base with num = %v\n", b.num)
	return "describer run successfully."
}

type container struct {
	base
	name string
}

func main() {
	cnt := container{
		base: base{
			num: 11,
		},
		name: "tanki",
	}
	fmt.Printf("cnt:{num: %v, name: %v\n", cnt.num, cnt.name)

	fmt.Println("another way to call num = ", cnt.base.num)

	fmt.Println(cnt.describe())

	type describer interface {
		describe() string
	}
	var d describer = cnt
	fmt.Println("describer = ", d.describe())
}
