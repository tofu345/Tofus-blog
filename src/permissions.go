package src

import (
	"log"
)

var permissions = map[string]string{
	"delete_post": "Can Delete Posts",
	"update_post": "Can Update Posts",
}

func init() {
	db.Exec("DELETE from permissions")
	registerPerms(permissions)
}

func registerPerms(perms map[string]string) {
	var perm Permission
	for k, v := range perms {
		perm = Permission{Name: k, Description: v}
		err := db.Create(&perm).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}

func userHasPerm(user User, perm string) bool {
	if _, exists := permissions[perm]; !exists {
		log.Printf("%v Permission does not exist", perm)
		return false
	}

	for _, v := range user.UserPerms {
		if v.Name == perm {
			return true
		}
	}

	return false
}
