package main

import "fmt"

func main() {
	// var firstName string = "Mechiuw"

	// lastName := "Diamonda"
	// age := 10
	// fmt.Println(firstName, lastName)
	// fmt.Printf("my age is : %d \n", age)

	// var (
	// 	bootCampName, bootCampAddress = "Enigma Camp", "Malang"
	// )
	// fmt.Printf("name : %s || address : %s", bootCampName, bootCampAddress)

	variables()
	multiple()
}

func variables() {
	username := "mechiuw"
	gmail := "matthewdpk@gmail.com"
	fmt.Printf("username : %s\n", username)
	fmt.Printf("gmail : %s\n", gmail)
}

func multiple() {
	var a, b, c, d string = "hey", "hi", "hello", "heyow"
	fmt.Println(a, b, c, d)

	var w, x, y, z = 10, 54, 12, 38
	fmt.Println(w, x, y, z)

	var number, text = 10, "this is ten"
	fmt.Println(number, text)

	one, two, three := 1, true, "woah"
	fmt.Println(one, two, three)

	var (
		x1        = 10
		x2 int    = 20
		x3 string = "example"
	)
	fmt.Println(x1, x2, x3)
}
