package main

import "fmt"

//This is a single line comment

/*
	This is a 
	multi line comments
*/

func main() {
	var radius float32
	fmt.Scanf("%f", &radius)

	var length float32
	length = 2 * 3.14 * radius

	fmt.Println("length:", length)
}

