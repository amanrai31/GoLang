package main

import (
	"fmt"
)

const GlobalName = "Rai"

func main() {
	var name string = "Aman" // var name ="Aman" - this is also acceptable[We can not change the variable's type]
	fmt.Println(name)
	fmt.Printf("Variable type of name: %T \n", name) //use Printf to get the type of variable.

	var age int = 23 // var age = 23 - we can declear thik this too.[We can not change the variable's type]
	fmt.Println(age)
	fmt.Printf("Variable type of age: %T \n", age)

	var emptyint int
	fmt.Println(emptyint)
	fmt.Printf("Variable type of emptyint: %T \n", emptyint)

	var emptystring string
	fmt.Println(emptystring) // This will print nothing
	fmt.Printf("Variable type of emptystring: %T \n", emptystring)

	anotherName := "Appu" // We can declear vaiables this this, But only allowed inside function.[We can not change the variable's type]
	anotherage := 24
	fmt.Println("\n Other name is:", anotherName, "and the age is:", anotherage)

	fmt.Println("Global name is:", GlobalName)
	fmt.Printf("Global type is: %T", GlobalName)

}
