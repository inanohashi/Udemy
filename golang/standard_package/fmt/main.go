package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("表示")

	//fmt表示
	fmt.Print("Hello\n")
	//改行
	fmt.Println("Hello!")

	//println系 各々の文字列は半角スペースで区切られ、文字列の最後に改行を追加する
	fmt.Println("Hello", "world!")
	fmt.Println("Hello", "world!")

	//Printf系 フォーマットを指定
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")

	//Sprint系 出力ではなくフォーマットした結果を文字列で返す。変数に代入時な
	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", "Hello")
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	//Fprint系 書き込み先指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")
}
