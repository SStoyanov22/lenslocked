package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name  string
	Bio   string
	Role  Role
	Items map[string]string
}

type Role struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	role := Role{
		Name: "Admin",
	}
	items := map[string]string{}

	items["key1"] = "value1"
	items["key2"] = "value2"
	user := User{
		Name:  "Stoyan Stoyanov",
		Bio:   `<script>alert("Haha, you have been hacked!")</script>`,
		Role:  role,
		Items: items,
	}

	fmt.Println(items["key1"])
	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
