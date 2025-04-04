package main

import "fmt"

func main() {

	var age int = 30
	var count int32 = 10

	fmt.Println(age, count)

	// int8:** Inteiro de 8 bits com sinal. Valores de -128 a 127.
	// int16:** Inteiro de 16 bits com sinal. Valores de -32768 a 32767.
	// int32:** Inteiro de 32 bits com sinal. Valores de -2147483648 a 2147483647.
	// int64:** Inteiro de 64 bits com sinal. Valores de -9223372036854775808 a 9223372036854775807.	
	var index int8 = 127
	var index2 int16 = 32767
	var index3 int32 = 2147483647
	var index4 int64 = 9223372036854775807
	fmt.Println(index, index2, index3, index4)
}