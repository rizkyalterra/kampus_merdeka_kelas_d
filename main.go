package main

import "fmt"

func main() {
	fmt.Println("stash")

	// interface
	var userInterface UserInterface

	user := User{"Alterra", 12}
	user.tambahUmur()

	userInterface = user
	userInterface.tambahUmur()

}

type User struct {
	Name string
	Umur int
}

func (user User) tambahUmur() {
	fmt.Println(user.Umur + 10)
}

type UserInterface interface {
	tambahUmur()
}
