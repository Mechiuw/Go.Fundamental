package main

import "fmt"

func main() {
	i := 20.2
	txt := "texts here"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}

// nah there are more than these, imma straight search for it if i need
// https://www.w3schools.com/go/go_formatting_verbs.php
