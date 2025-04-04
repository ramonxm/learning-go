package main

import "fmt"

func main() {
	// String
	var name string = "John"
	fmt.Println(name)
	
	// Multiline string
	var multiline string = `
	Hello World
	How are you?
	`
	fmt.Println(multiline)

	// String length
	fmt.Println(len(name))
	
	// String concatenation
	fmt.Println(name + " " + "Doe")
	
	// String interpolation
	fmt.Printf("My name is %s and I am %d years old\n", name, 20)
	
	// String to rune
	var runeName []rune = []rune(name)
	fmt.Println(runeName)
	
	// Rune to string
	var stringName string = string(runeName)
	fmt.Println(stringName)
	
	// String to byte
	var byteName []byte = []byte(name)
	fmt.Println(byteName)
	
	// Byte to string
	var stringName2 string = string(byteName)
	fmt.Println(stringName2)		
}