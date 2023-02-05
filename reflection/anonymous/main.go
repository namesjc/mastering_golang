package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Manager struct {
	User
	Title string
}

func main() {
	m := Manager{
		User:  User{1, "ok", 12},
		Title: "Manager",
	}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}
