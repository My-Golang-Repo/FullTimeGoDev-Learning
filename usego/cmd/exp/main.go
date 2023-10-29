package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  template.HTML
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Smith",
		Bio:  `<script>alert("Haha")</script>`,
	}

	if err := t.Execute(os.Stdout, user); err != nil {
		panic(err)
	}
}
