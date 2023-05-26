package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//構造体からJSONテキストへの変換

type A struct{}

type User struct {
	ID      int       `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Email   string    `json:"email,omitempty"`
	Created time.Time `json:"created,omitempty"`
	A       *A        `json:"A,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)
	u2 := new(User)

	//Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}
