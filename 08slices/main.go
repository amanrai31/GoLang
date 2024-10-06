package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SLICES")

	slice1 := []int{105, 200, 186, 875, 321}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	slice1 = append(slice1, 333, 222) // Appending values.
	fmt.Println(slice1)
	fmt.Println(slice1[0:4]) //Last element is not included.
	fmt.Println(slice1[0:])
	fmt.Println(slice1[1:6])
	sort.Ints(slice1) // Sorting
	fmt.Println(slice1)
}
