package main

import "fmt"

func main() {
	fmt.Println(hey(5, 5))
	fmt.Println(hey2(5, 5))
	fmt.Println(hey3(10, 5, "hey", "there"))
}

//regular returns
func hey(x int, y int) int {
	return x + y
}

//defined returns
func hey2(x int, y int) (result int) {
	result = x + y
	return
}

// multiple returns
func hey3(x int, y int, text string, text2 string) (result int, message string, concated string) {
	result = x + y
	message = text
	concated = text + text2
	return
}
