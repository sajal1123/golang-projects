package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cannot work with 42!")
	}
	return arg + 100, nil
}

type argError struct {
	n    int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.n, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Cannot work with it!"}
	}
	return arg + 100, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1() does not work!", r, e)
		} else {
			fmt.Println("f1() works!", r, e)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2() does not work!", r, e)
		} else {
			fmt.Println("f2() works!", r, e)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.n)
		fmt.Println(ae.prob)
	}
}
