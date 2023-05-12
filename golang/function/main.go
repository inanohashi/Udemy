package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := ReturnFunc()

	f()
	CallFunction(func() {
		fmt.Println("I'm a function")
	})

	fun := Later()
	fmt.Println(fun("Hello"))
	fmt.Println(fun("My"))
	fmt.Println(fun("name"))
	fmt.Println(fun("is"))
	fmt.Println(fun("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

}
