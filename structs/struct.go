package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{Name: "Aditya", Email: "aditya.dev2516@gmail.com", Age: 22}
	fmt.Println("User 1:", u1)
}
