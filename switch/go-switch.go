package main

import "fmt"

func main() {
	switches()
}

func switches() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		switch i {
		case 1:
			fmt.Println("fish")
		case 2:
			fmt.Println("bush")
		}
	}
}
