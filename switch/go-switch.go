package main

import "fmt"

func main() {
	switches()
}

func switches() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 1:
			fmt.Println("fish")
		case 2:
			fmt.Println("bush")
		case 3:
			fmt.Println("fishbush")
		case 4:
			fmt.Println("bushfish")
		case 5:
			fmt.Println("fishesbushes")
		}
	}
}
