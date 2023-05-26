package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "c", "b"}

	fmt.Println(i, s)
	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s)
}
