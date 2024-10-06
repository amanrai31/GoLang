package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	var ptr *int
	fmt.Println("Value of unassigned pointer=", ptr)

	num1 := 69
	ptr1 := &num1
	fmt.Println(ptr1)
	fmt.Println(*ptr1)

	fmt.Println(*ptr1 + 5)

}
