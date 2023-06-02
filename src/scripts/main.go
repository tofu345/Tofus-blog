package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tofs-blog/src"

	"gorm.io/gorm"
)

var db *gorm.DB
var r = bufio.NewReader(os.Stdin)
var scripts = map[string]func(){
	"create_admin": createAdmin,
}

func init() {
	db = src.GetDB()
}

func main() {
	function, exists := scripts[getUserInput("> Run Script: ")]
	if !exists {
		fmt.Println("! Script Not Found")
		return
	}

	function()
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := r.ReadString('\n')
	if text == "" {
		fmt.Println("! This field is required")
		return getUserInput(prompt)
	}
	return strings.TrimSpace(text)
}
