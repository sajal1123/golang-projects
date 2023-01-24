package main

import "fmt"

//func main() {
//	const r = 10
//	i := 0
//	for i < r {
//		fmt.Println(i)
//		i++
//	}
//	const s = 10
//	j := 0
//	fmt.Println("Value of i is", i)
//	for j < s {
//		i := 10
//		for i > 0 {
//			if j >= i {
//				fmt.Print("*")
//			} else {
//				fmt.Print(' ')
//			}
//			i--
//		}
//		fmt.Println('\n')
//		j++
//	}
//}

func main() {
	const r = 10
	fmt.Println("First loop")
	for i := 0; i < r; i++ {
		fmt.Println(i)
	}
	j := 21
	fmt.Println("Second loop")
	for j > r {
		fmt.Println(j)
		j--
	}
	k := 1
	fmt.Println("Third Loop")
	for {
		fmt.Print(k + ' ')
		if k > 100 {
			break
		}
		k++
	}
}
