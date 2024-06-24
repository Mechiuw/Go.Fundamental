package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr2)

	arr3 := [...]int{4, 3, 2, 1}
	fmt.Println(arr3)

	arr4 := [...]string{"bmw", "toyota", "honda"}
	fmt.Println(arr4[1])
	fmt.Println(arr4[2])

	arr4[2] = "ford"
	fmt.Println(arr4[2])

	initial()
}

func initial() {
	ar := [5]int{1, 2}
	fmt.Println(ar)

	arx := [3]string{"hey", "there"}
	fmt.Println(arx)

	specific()
}

func specific() {
	arr1 := [5]int{1: 10, 2: 40, 0: 3}
	fmt.Println(arr1)

	lengthArr()
}

func lengthArr() {
	arx := [...]int{1, 23, 45, 6, 6, 74, 3}
	arrLength := len(arx)
	fmt.Println(arrLength)
}
