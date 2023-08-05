package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/tofu345/Tofus-blog/src"

	"golang.org/x/term"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	r       = bufio.NewReader(os.Stdin)
	scripts = []Script{
		{"create_admin", "Create Admin User", createAdmin},
		{"give_admin", "Give User Admin Permissions to User", giveAdminPerm},
	}
	loggedInAdmin src.User
)

type Script struct {
	name        string
	description string
	function    func()
}

func init() {
	db = src.GetDB()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("> This must be ran from project root")
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("? 'help' to view list of commands")

	for {
		input := getUserInput("> ")

		switch input {
		case "":
			continue
		case "help":
			fmt.Println("list\tlist all commands")
			fmt.Println("exit\tquit")
		case "list":
			if len(scripts) == 0 {
				fmt.Println("! There are no scripts")
				return
			}

			for _, v := range scripts {
				fmt.Printf("%v\t%v\n", v.name, v.description)
			}
		case "exit":
			return
		default:
			script, err := getScript(input)
			if err != nil {
				fmt.Printf("! %v\n", err)
				continue
			}

			script.function()
		}
	}
}

func getScript(name string) (Script, error) {
	for _, v := range scripts {
		if v.name == name {
			return v, nil
		}
	}

	if strings.HasPrefix(name, "ex") {
		os.Exit(0)
	}

	return Script{}, fmt.Errorf("Script '%v' not found", name)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := r.ReadString('\n')
	if text == "exit" {
		os.Exit(0)
	} else if text == "" {
		fmt.Println("! This field is required")
		return getUserInput(prompt)
	}
	return strings.TrimSpace(text)
}

func adminLogin() (src.User, error) {
	if loggedInAdmin.ID != 0 {
		return loggedInAdmin, nil
	}

	fmt.Println("! Admin Login Required")

	email := getUserInput("> Admin Email: ")

	admin := src.User{}
	err := db.Where("email = ?", email).Find(&admin).Error
	if err != nil {
		return src.User{}, err
	}

	if admin.ID == 0 {
		fmt.Printf("! No user found with email '%v'\n", email)
		return adminLogin()
	}

	fmt.Print("> Admin Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return src.User{}, err
	}
	password := string(bytePassword)
	fmt.Println()

	if !src.CheckPasswordHash(password, admin.Password) {
		fmt.Println("Passwords do not match")
		return adminLogin()
	}

	if !admin.IsAdmin {
		fmt.Printf("! %v does not have admin permissions\n", admin.Username)
		return adminLogin()
	}

	loggedInAdmin = admin

	return admin, nil
}

func Fatal(err error) {
	fmt.Printf("! %v\n", err)
	os.Exit(0)
}

func getPassword() string {
	fmt.Print("> Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		Fatal(err)
	}
	password := string(bytePassword)
	fmt.Println()

	fmt.Print("> Retype Password: ")
	bytePassword2, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		Fatal(err)
	}
	password2 := string(bytePassword2)
	fmt.Println()

	if password != password2 {
		fmt.Println("! Passwords do not match")
		return getPassword()
	}

	password, err = src.HashPassword(password)
	if err != nil {
		Fatal(err)
	}

	return password
}

// Scripts
func createAdmin() {
	admins := []src.User{}
	err := db.Where("is_admin <> ?", true).Find(&admins).Error
	if err != nil {
		Fatal(err)
	}

	if len(admins) != 0 {
		_, err := adminLogin()
		if err != nil {
			Fatal(err)
		}
	}

	fmt.Println(">> Create Admin user")

	username := getUserInput("> Username: ")
	email := getUserInput("> Email: ")
	password := getPassword()

	user := src.User{
		FirstName: username,
		Username:  username,
		Password:  password,
		Email:     email,
		IsAdmin:   true,
	}

	err = db.Create(&user).Error
	if err != nil {
		Fatal(err)
		return
	}

	fmt.Println("! Admin User Created Successfully")
}

func giveAdminPerm() {
	admin, err := adminLogin()
	if err != nil {
		Fatal(err)
	}

	fmt.Println(">> Give User Admin Permissions")

	user_email := getUserInput("> Enter User Email: ")
	var user src.User
	err = db.First(&user, "email = ?", user_email).Error
	if err != nil {
		Fatal(err)
	}

	input := getUserInput(fmt.Sprintf("> Give %v admin permissions? (y/n): ", user.Username))
	if input != "y" {
		fmt.Println("! Aborted")
		return
	}

	user.IsAdmin = true
	err = db.Save(&admin).Error
	if err != nil {
		fmt.Println("! Error Saving User Data")
		return
	}

	fmt.Printf("! %v is now an admin\n", user.Username)
}
