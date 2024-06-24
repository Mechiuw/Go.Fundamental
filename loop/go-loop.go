package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "pear", "kiwi", "dragon fruit", "watermelon"}
	for idx, value := range fruits {
		if idx == 0 {
			continue
		}
		fmt.Println(idx, value)
	}
}
