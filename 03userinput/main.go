package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your fav dish=")
	input, _ := reader.ReadString('\n')

	fmt.Println("Your fav dish is =>", input)

	var rating int
	fmt.Println("How you rate your fav dish", input)
	fmt.Scan(&rating)
	fmt.Print(input, "rating is =", rating)

	var num1 int
	var num2 int
	fmt.Println("\n Enter numbers=")
	fmt.Scan(&num1, &num2)
	fmt.Println(num1, num2)

	var str1 string
	fmt.Println("Enter string=")
	fmt.Scanln(&str1)
	fmt.Println(str1)
}
