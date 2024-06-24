package main

import "fmt"

func main() {
	printFunction()
	printfFunction()
	printlnFunction()
}

func printFunction() {
	i := "print() function"
	fmt.Print(i, "\n")
}
func printfFunction() {
	j := "printf() function"
	fmt.Printf("%s\n", j)
}
func printlnFunction() {
	k := "println() function"
	fmt.Println(k)
}
