package main

import (
	"fmt"
)

const MyuserName = "Yash Srivastava" // global type constant

func main() {
	fmt.Println("This is about variables and constants in GOlang")

	// usual way of declaring a variable
	var userName string = "Yash Srivastva"
	fmt.Println(userName)

	// second way the implicit way
	var username1 = "Yash Srivastava"
	fmt.Println(username1)

	// third way the no var style
	userName2 := "Yash Srivastava"
	fmt.Println(userName2)

	// now declaring the constant.
	const myUserName = "Yash Srivastava" // this type of constant is non public one
	fmt.Println(myUserName)
}
