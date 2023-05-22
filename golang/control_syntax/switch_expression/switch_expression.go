package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3or 4")
	default:
		fmt.Println("I don't Know")
	}

	// 変数を評価したい場合は変数の局所性を高められるため、こちらが望ましい
	switch n := 1; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}

	n = 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n <4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}

	/* 列挙型を混在させるとエラーが出る
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	default:
		fmt.Println("I don't Know")
	}
	*/
}
