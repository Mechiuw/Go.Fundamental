package main

import "fmt"

func main() {
	var person1 Person

	indx := []int{0, 1, 2, 3, 4, 5}

	person1.name = "mason"
	person1.age = 34
	person1.job = "usmc"
	person1.salary = 1200

	for idx := range indx {
		fmt.Println(idx, person1.name, person1.age, person1.job, person1.salary)
	}

}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}
