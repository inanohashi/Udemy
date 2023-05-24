package main

import (
	"fmt"
	"time"
)

func main() {
	// 時間の生成
	// 今の時間
	t := time.Now()
	fmt.Println(t)

	// 指定した時間を生成
	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.YearDay())

	// 時間の感覚を表現する
	// time.Duration型を返す
	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)

	// time.Duration方を文字列から生成する
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	// 現在時刻の２分３０秒後を表すtime.Time型の取得
	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)

	// 時刻の差分を取得
	t5 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	t6 := time.Now()
	fmt.Println(t6)

	//t5 - t6
	d2 := t5.Sub(t6)
	fmt.Println(d2)

	// 時刻を比較してbool値を返す
	fmt.Println(t6.Before(t5))
	fmt.Println(t6.After(t5))
	fmt.Println(t5.Before(t6))
	fmt.Println(t5.After(t6))

	// 指定時間のスリープ
	// ５秒間停止
	time.Sleep(5 * time.Second)
	fmt.Println("５秒停止後表示")
}
