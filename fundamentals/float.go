package main

import "fmt"

func main() {

	var num float32 = 30.5

	fmt.Println(num)

	var pi float64 = 3.14
	var radius float64 = 10.5

	var area float64 = pi * radius * radius
	fmt.Println(area)

	// Float32 + Float64 resultaria erro (mismatched types float32 and float64)
}