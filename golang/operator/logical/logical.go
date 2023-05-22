package main

import "fmt"

func main() {
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	fmt.Println(!true)
	fmt.Println(!false)
}
