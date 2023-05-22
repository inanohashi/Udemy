package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000
	fmt.Println(sl)

	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2)

	sl3 := []int{100, 200}
	sl4 := make([]int, 2, 5)
	n := copy(sl4, sl3)
	fmt.Println(n, sl4)

}
