package main

import "fmt"

func main() {
	// var a = map[string]string{"brand": "ford", "model": "mustang", "year": "1964"}
	// fmt.Println(a)

	var b = make(map[int]string)
	b[1] = "teacher"
	b[2] = "techinician"
	b[3] = "engineer"
	b[4] = "judge"
	b[5] = "entrepeneur"

	for i := 1; i < len(b); i++ {
		fmt.Println(i, b[i])
	}

	//empty map
	var c = make(map[string]string)
	var d map[string]string

	fmt.Println(c)
	fmt.Println(d)
}
