package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//入力全体を読み込む
	f, _ := os.Open("foo.txt")
	//ioで代替
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	//ファイルに書き込み
	//osで代替
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}

}
