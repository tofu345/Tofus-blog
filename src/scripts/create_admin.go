package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"tofs-blog/src"

	"golang.org/x/term"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = src.GetDB()
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("> Create Admin user")
	username := getUserInput(r, "> Username: ")
	email := getUserInput(r, "> Email: ")

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
		Username:    username,
		Password:    password,
		Email:       email,
		AccessToken: src.GenerateToken(email), // Leaving null causes issues with unique property
		IsAdmin:     true,
	}

	err = db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("! Admin User Created Successfully")
}

func getUserInput(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := r.ReadString('\n')
	if text == "" {
		fmt.Println("! This field is required")
		return getUserInput(r, prompt)
	}
	return strings.TrimSpace(text)
}
