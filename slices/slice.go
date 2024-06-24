package main

import "fmt"

func main() {
	slice_name := []int{1, 2, 3, 4, 5} // is like a List<Integer> from java
	fmt.Println(slice_name)
	fmt.Println(cap(slice_name))
	fmt.Println(len(slice_name))

	my_slice := []int{}
	fmt.Println(my_slice)

	another := []int{10, 11, 12, 13, 14, 15}
	slice_another := another[2:6]
	fmt.Println(slice_another)
}
