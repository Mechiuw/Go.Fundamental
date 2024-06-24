package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arx := []string{"volvo", "bmw", "vw", "lambo", "pagani"}
	passArgs("passed argument ")
	passArr(arr, arx)
}

func passArgs(arg string) {
	fmt.Print(arg, "\n")
}

func passArr(arr []int, arx []string) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Println(arr[i], "is even")
		} else {
			fmt.Println(arr[i], "is odd")
		}
	}

	for idx, val := range arx {
		fmt.Println(idx, val)
	}
}
