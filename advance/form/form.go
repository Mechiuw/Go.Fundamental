package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	formInput()
}

func formInput() {
	var age int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=========FORM===========")
	fmt.Println("						 ")

	fmt.Print("type your first name : ")
	scanner.Scan()
	firstName := scanner.Text()

	fmt.Print("type your last name : ")
	scanner.Scan()
	lastName := scanner.Text()

	fmt.Print("type your age : ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())

	fmt.Print("type your address : ")
	scanner.Scan()
	address := scanner.Text()

	fmt.Print("type your email : ")
	scanner.Scan()
	email := scanner.Text()

	formMaker(firstName, lastName, age, address, email)
}

func formMaker(firstName string, lastname string, age int, address string, email string) {
	fmt.Println("NAME 		:", firstName, lastname)
	fmt.Println("AGE 		:", age)
	fmt.Println("ADDRESS 	:", address)
	fmt.Println("EMAIL 		:", email)
}
