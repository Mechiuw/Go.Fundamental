package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(concat("matthew ", "diamonda ", "posma ", "kesaulya "))
}

func sum(num ...int) int {
	numbrs := 0
	for _, val := range num {
		numbrs += val
	}
	return numbrs
}

func concat(name ...string) string {
	concated := ""
	for _, val := range name {
		concated += val
	}
	return concated
}
