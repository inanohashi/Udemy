package main

import (
	"database/sql"
	//"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// テーブルの更新
	/*
		cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
		_, err := Db.Exec(cmd, "tarou", 20)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// 特定のデータを取得
	/*
		cmd := "SElECT * FROM persons WHERE age = $1"
		//QueryRow 1レコード取得
		row := Db.QueryRow(cmd, 20)
		var p Person
		err := row.Scan(&p.Name, &p.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No row")
			} else {
				log.Println(err)
			}
		}
		fmt.Println(p.Name, p.Age)
	*/

	// 複数のデータを取得
	/*
		cmd := "SELECT * FROM persons"
		// Queryは条件に合うものを全て取得
		rows, _ := Db.Query(cmd)
		defer rows.Close()
		var pp []Person
		for rows.Next() {
			var p Person
			err := rows.Scan(&p.Name, &p.Age)
			if err != nil {
				log.Fatalln(err)
			}
			pp = append(pp, p)
		}

		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	*/

	// データの更新
	/*
		cmd := "UPDATE persons SET age = $1 WHERE name = $2"
		_, err := Db.Exec(cmd, 25, "tarou")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// データの削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "tarou")
	if err != nil {
		log.Fatalln(err)
	}

}
