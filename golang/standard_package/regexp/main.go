package main

import (
	"fmt"
	"regexp"
)

func main() {
	//Goの正規表現の基本 Regexp.MatchString
	/*
		match, _ := regexp.MatchString("A", "ABC")
		fmt.Println(match)

		//Compile
		re1, _ := regexp.Compile("A")
		match = re1.MatchString("ABC")
		fmt.Println(match)

		//MustCompile
		re2 := regexp.MustCompile("A")
		match = re2.MatchString("ABC")
		fmt.Println(match)
	*/

	/*
		regexp.MustCompile("\\d")
		regexp.MustCompile(`\d`)
	*/

	//正規表現のフラグ
	/*
		re3 := regexp.MustCompile(`(?i)abc`)
		match := re3.MatchString("ABC")
		fmt.Println(match)
	*/

	/*
		//幅を持たない正規表現のパターン

		re4 := regexp.MustCompile(`^ABC$`)
		match := re4.MatchString("ABC")
		fmt.Println(match)

		match = re4.MatchString(" ABC ")
		fmt.Println(match)


		//繰り返しを表す正規表現
		re6 := regexp.MustCompile("a+b*")
		fmt.Println(re6.MatchString("ab"))
		fmt.Println(re6.MatchString("a"))
		fmt.Println(re6.MatchString("aaaabbbbbbbbbbb"))
		fmt.Println(re6.MatchString("b"))

		//正規表現の文字クラス
		re8 := regexp.MustCompile(`[XYZ]`)
		fmt.Println(re8.MatchString("Y"))

		re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
		fmt.Println(re9.MatchString("ABC"))
		fmt.Println(re9.MatchString("abcdefg"))

		re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
		fmt.Println(re10.MatchString("ABC"))
		fmt.Println(re10.MatchString("あ"))

		//正規表現のグループ
		re11 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
		fmt.Println(re11.MatchString("abcxyz"))
		fmt.Println(re11.MatchString("ABCXYZ"))
		fmt.Println(re11.MatchString("ABCxyz"))
		fmt.Println(re11.MatchString("ABcabc"))


		re := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))

		//正規表現にマッチした文字列の取得
		//FindString
		fs1 := re.FindString("AAAABCXYZZZ")
		fmt.Println(fs1)
		fs2 := re.FindAllString("ABCXVZABCXVZ", -1)
		fmt.Println(fs2)

		//正規表現のグループによるサブマッチ
		re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
		s := `
				0123-456-7889
				111-222-333
				556-787-899
			`
		ms := re1.FindAllStringSubmatch(s, -1)
		fmt.Println(ms)
		for _, v := range ms {
			fmt.Println(v)
		}

		fmt.Println(ms[0][0])
		fmt.Println(ms[0][1])
		fmt.Println(ms[0][2])
	*/

	//正規表現による文字列の置換
	re1 := regexp.MustCompile(`\s+`)
	fmt.Println(re1.ReplaceAllLiteralString("AAA BBB CCC", ","))

	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllLiteralString("私はGolangを使用する、プログラマー", ","))

	//正規表現による文字列の分割
	re3 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re3.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("aaaaaa    bbbb  cccccccc", -1))
}
