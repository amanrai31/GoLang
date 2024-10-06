package main

import "fmt"

func main() {
	fmt.Println("Array")

	array1 := []string{"aman", "abc", "xyz"} //4th element is empty=> " "
	fmt.Println("Array1 =>", array1)
	fmt.Println(len(array1)) // len(array)

	var array2 [3]string

	array2[0] = "1"
	array2[1] = "2"
	array2[2] = "3"
	fmt.Println(array2)

	var array3 [5]int //Empty array of type int => default value is 0
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5, 6, 7} //Go infers the sie.
	fmt.Println(len(array4), "=>", array4)

	array5 := [2][3]int{
		{1, 2, 3}, {4, 5, 6},
	}

	fmt.Println(array5)

	array6 := [][]string{
		{"a", "b", "c"}, {"d", "e", "f"},
	}
	fmt.Println(array6)

	fmt.Printf("type of array1 %T \n", array1)
}
