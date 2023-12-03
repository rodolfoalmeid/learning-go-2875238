package main

import (
	"fmt"
)

func main() {
	var myint int = 45
	fmt.Println("This is my value", myint)
	fmt.Printf("The Variable Type is %T\n", myint)

	mystring := "This my string"
	fmt.Println(mystring)
	fmt.Printf("This is the variable type %T\n", mystring)
}
