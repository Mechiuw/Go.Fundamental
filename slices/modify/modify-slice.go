package main

import "fmt"

func main() {
	modifySlice()
}

func modifySlice() {
	prices := []int{10, 20, 30, 40, 50}
	pay := 30

	for i := 0; i < len(prices); i++ {
		if pay < prices[i] {
			fmt.Println("above ", pay)
		} else if pay > prices[i] {
			fmt.Println("buy for ", prices[i], "!")
		} else {
			fmt.Println("i'll pass for ", prices[i])
		}
	}
}
