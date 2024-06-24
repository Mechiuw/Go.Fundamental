package main

import "fmt"

const PI = 3
const A = 21
const B = 21 //UNTYPED CONSTANTS

func main() {
	d := 20
	result := PI * d
	fmt.Println("result: ", result)
	two()
}

func two() {
	result := A * PI
	fmt.Println("two: ", result)
	three()
}

func three() {
	const UNCHANGEABLE = 20 // unchangeable and constant unlike outside function scope
	fmt.Println("unchangeable const: ", UNCHANGEABLE)
}
