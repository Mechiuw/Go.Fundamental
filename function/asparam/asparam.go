package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNum := Filter(numbers, isEven)
	oddNum := Filter(numbers, isOdd)
	fmt.Println("even numbers : ", evenNum)
	fmt.Println("odd numbers : ", oddNum)
}

func Filter(arr []int, cond func(int) bool) []int {
	var result []int
	for _, val := range arr {
		if cond(val) {
			result = append(result, val)
		}
	}
	return result
}

func isOdd(i int) bool {
	return i%2 == 0
}

func isEven(i int) bool {
	return i%2 != 0
}
