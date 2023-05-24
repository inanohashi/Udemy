package main

import "fmt"

func main() {
	var ch1 chan int
	// 受信専用 var che <-chan int
	// 送信専用 var ch3 chan<- int

	ch1 = make(chan int)

	fmt.Println(cap(ch1))

	ch2 := make(chan int, 5)
	fmt.Println(cap(ch2))

	ch2 <- 1
	fmt.Println(len(ch2))

	ch2 <- 2
	ch2 <- 3
	fmt.Println("len", len(ch2))

	i := <-ch2
	fmt.Println(i)
	fmt.Println("len", len(ch2))

	i2 := <-ch2
	fmt.Println(i2)
	fmt.Println("len", len(ch2))

	fmt.Println(<-ch2)
	fmt.Println("len", len(ch2))
}
