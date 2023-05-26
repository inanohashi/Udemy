package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	//MD5ハッシュ値を生成
	//任意の文字列からMD5ハッシュ値を生成する処理例
	h := md5.New()
	io.WriteString(h, "ABCDE")
	fmt.Println(h.Sum(nil))

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}
