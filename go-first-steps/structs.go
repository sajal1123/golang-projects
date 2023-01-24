package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p
}

func main() {

	fmt.Println(person{"Sajal", 24})
	fmt.Println(person{name: "Messi", age: 35})
	fmt.Println(person{name: "Suarez"})
	fmt.Println(newPerson("Pedri"))
	fmt.Println(&person{"iniesta", 36})

	s := person{name: "puyol", age: 38}
	fmt.Println("s = ", s)
	sp := &s
	fmt.Println("sp = ", sp)

	fmt.Println("sp.age = ", sp.age)
	sp.age = 11
	fmt.Println("sp.age = 11\nsp.age = ", sp.age)

	fmt.Println("s.age = ", s.age)
	s.age = 111
	fmt.Println("s.age = 111\ns.age = ", s.age)

}
