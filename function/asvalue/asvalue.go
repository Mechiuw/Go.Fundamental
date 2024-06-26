package main

import "fmt"

func main() {
	var summarize func(int, int) int = sum
	fmt.Println(summarize(5, 5))
}

func sum(a, b int) int {
	return a + b
}
