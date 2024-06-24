package main

import "fmt"

func main() {
	testCount(1)
}

func testCount(count int) int {
	if count == 11 {
		return 0
	}
	fmt.Println(count)
	return testCount(count + 1)
}
