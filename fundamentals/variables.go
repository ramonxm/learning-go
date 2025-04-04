package main

import "fmt"

func main() {
	// Variable declaration
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println(c)

	// Short hand declaration
	d := 10
	e := 20
	f := d + e
	fmt.Println(f)

	// Multiple declaration
	var g, h, i int = 10, 20, 30
	fmt.Println(g, h, i)
	
	// Multiple declaration with different types
	var j, k, l = 10, "Hello", true
	fmt.Println(j, k, l)
	
	// Multiple declaration with different types using short hand declaration
	m, n, o := 10, "Hello", true
	fmt.Println(m, n, o)
}