package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	var s string = "100"

	i2, _ := strconv.Atoi(s)
	fmt.Println(i2)
	fmt.Printf("i = %T\n", i2)

	i3, err := strconv.Atoi("A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i3)
	fmt.Printf("i = %T\n", i3)

	s2 := strconv.Itoa(i)
	fmt.Println(s2)
	fmt.Printf("s = %T\n", s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
