package main

import "fmt"

type T struct {
	// 埋め込みの場合は、[User User]を[User]として型名を省略してフィールドのみで記入ができる
	User User
}

type User struct {
	Name string
	Age  int
}

func (u *User) setName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	// 埋め込みで型名を省略した場合は、直接内側のフィールドにアクセスできる
	// fmt.Println(t.Name)

	t.User.setName()
	fmt.Println(t)
}
