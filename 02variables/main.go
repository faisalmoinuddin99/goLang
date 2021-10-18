package main

import "fmt"

func main() {
	// In go lang, everthing is an type
	/*
		1. To create a variable we uses "var" keyword
		2. Variabl type should be known in advance
	*/

	var message string = "hello"
	fmt.Println(message)                            // hello
	fmt.Printf("variable is of type: %T ", message) // variable is of type: string

	//boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)                           // false
	fmt.Printf("variable is of type: %T", isLoggedIn) // variable is of type: bool

	//default value
	var someRandomVariable int
	fmt.Print("\n")
	fmt.Println(someRandomVariable) // 0

	/*
		# If we notice the above 3 lines
		# The variable is only decllared but not initialize
		# So the default value is 0
	*/

	// IMPLICIT TYPE
	var website = "www.google.com"
	fmt.Println(website)

	/*
		# if we run the above code it will run successfully
		# but we know that in go lang variable type should be declared
		# but here we didnt declare the type of variabel
		# so for this, lexer comes into the picture and take care of that
		# Base on value we put into variable, Lexer will automatically decide the type of variable
	*/

	/*
		NOTE: But point to remember, if we try to change the valye from one type to another
		go lang will throws an error
	*/
	// website = 3 ; error

}
