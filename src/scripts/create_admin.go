package main

import (
	"fmt"
	"log"
	"syscall"
	"tofs-blog/src"

	"golang.org/x/term"
)

func createAdmin() {
	fmt.Println("> Create Admin user")
	username := getUserInput("> Username: ")
	email := getUserInput("> Email: ")

	fmt.Print("> Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	password := string(bytePassword)
	fmt.Println()

	password, err = src.HashPassword(password)
	if err != nil {
		log.Fatal(err)
		return
	}

	user := src.User{
		FirstName:   username,
		Username:    username,
		Password:    password,
		Email:       email,
		AccessToken: email, // Leaving null causes issues with unique property
		IsAdmin:     true,
	}

	err = db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("! Admin User Created Successfully")
}
