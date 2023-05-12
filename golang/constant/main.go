package main

import "fmt"

const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C = "A"
	D
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D)
	fmt.Println(c0, c1, c2)
}
