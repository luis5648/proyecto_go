package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Println("Introduzca el valor de a: ")
	fmt.Scanln(&a)

	for i := 0; i < 10; i++ {
		fmt.Println(a, ", ")

		fmt.Println("Hello world")

		a++

	}
	String1()
}
