package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var age int

	fmt.Println("==================================")

	scanner := bufio.NewScanner(os.Stdin) // init scanner
	fmt.Println("Data diri trainee enigma camp")

	fmt.Print("masukkan nama : ")
	scanner.Scan()
	fullname := scanner.Text()

	fmt.Print("masukkan umur : ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())

	fmt.Print("masukkan alamat : ")
	scanner.Scan()
	address := scanner.Text()

	fmt.Printf("%s | %d | %s", fullname, age, address)
}
