package src

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func postList(w http.ResponseWriter, r *http.Request) {
	postList := []Post{}
	err := db.Find(&postList).Error
	// ErrRecordNotFound if no posts
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "Post list", "data": postList})
}

func postDetail(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		JSONError(w, err)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "Post detail", "data": post})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	user, err := userFromBearer(w, r)
	if err != nil {
		JSONError(w, err)
		return
	}

	var post Post
	err = JSONDecode(r, &post)
	if err != nil {
		JSONError(w, err)
		return
	}

	post.AuthorID = user.ID
	post.AuthorName = user.Username
	maxChars := 60
	post.Slug = post.Title
	if maxChars < len(post.Slug) {
		post.Slug = post.Slug[:strings.LastIndex(post.Slug[:maxChars], " ")]
	}

	post.Slug = slug.Make(post.Slug)
	err_map := post.Errors()
	if len(err_map) != 0 {
		JSONError(w, err_map)
		return
	}

	err = db.Create(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "Post created", "data": post})
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		JSONError(w, err)
		return
	}

	user, err := userFromBearer(w, r)
	if err != nil {
		JSONError(w, err)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	if !userHasPerm(user, "delete_post") && post.AuthorID != user.ID {
		JSONError(w, Unauthorized)
		return
	}

	err = db.Delete(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "Post deleted"})
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	user, err := userFromBearer(w, r)
	if err != nil {
		JSONError(w, err)
		return
	}

	id, err := idParam(r)
	if err != nil {
		JSONError(w, err)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	if !userHasPerm(user, "delete_post") && post.AuthorID != user.ID {
		JSONError(w, Unauthorized)
		return
	}

	oldTitle := post.Title
	err = JSONDecode(r, &post)
	if err != nil {
		JSONError(w, err)
		return
	}

	post.ID = id // in case id is passed
	if oldTitle != post.Title {
		post.Slug = slug.Make(post.Title) // Regenerate slug
	}

	errs := post.Errors()
	if len(errs) != 0 {
		JSONError(w, errs)
		return
	}

	err = db.Save(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "Post updated", "data": post})
}

func createComment(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		JSONError(w, err)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	var comment Comment
	err = JSONDecode(r, &comment)
	if err != nil {
		JSONError(w, err)
		return
	}

	post.Comments = append(post.Comments, comment)
	err = db.Save(&post).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "Comment created", "data": post})
}

func userSignup(w http.ResponseWriter, r *http.Request) {
	var userData map[string]string
	err := JSONDecode(r, &userData)
	if err != nil {
		JSONError(w, err)
		return
	}

	user := User{
		Username:  userData["username"],
		Email:     userData["email"],
		Password:  userData["password"],
		FirstName: userData["first_name"],
		LastName:  userData["last_name"],
	}

	err_map := user.Errors()
	if len(err_map) != 0 {
		JSONError(w, err_map)
		return
	}

	var pswd string
	pswd, err = HashPassword(user.Password)
	if err != nil {
		JSONError(w, err)
		return
	}
	user.Password = pswd

	err = db.Create(&user).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "User created", "data": user})
}

func userList(w http.ResponseWriter, r *http.Request) {
	user, err := userFromBearer(w, r)
	if err != nil {
		JSONError(w, err)
		return
	}

	if !user.IsAdmin {
		JSONError(w, Unauthorized)
		return
	}

	users := []User{}
	err = db.Find(&users).Preload("Permissions").Error
	if err != nil {
		JSONError(w, err)
		return
	}

	Response(w, 200, dict{"message": "User list", "data": users})
}

func getAccessToken(w http.ResponseWriter, r *http.Request) {
	userData := map[string]string{}
	err := JSONDecode(r, &userData)
	if err != nil {
		JSONError(w, err)
		return
	}

	errors_map := map[string]string{}
	if _, exists := userData["email"]; !exists {
		errors_map["email"] = "This field is required"
	}
	if _, exists := userData["password"]; !exists {
		errors_map["password"] = "This field is required"
	}
	if len(errors_map) != 0 {
		JSONError(w, errors_map)
		return
	}

	var user User
	err = db.First(&user, "email = ?", userData["email"]).Error
	if err != nil {
		JSONError(w, err)
		return
	}

	if !CheckPasswordHash(userData["password"], user.Password) {
		JSONError(w, LoginError)
		return
	}

	access, err := newAccessToken(user.Username)
	if err != nil {
		JSONError(w, "Error Generating Token")
		return
	}

	refresh, err := newRefreshToken(user.Username)
	if err != nil {
		JSONError(w, "Error Generating Token")
		return
	}

	Response(w, 200, dict{"access": access, "refresh": refresh})
}

func refreshAccessToken(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := JSONDecode(r, &data)
	if err != nil {
		JSONError(w, err)
		return
	}

	if _, exists := data["token"]; !exists {
		JSONError(w, dict{"token": "This field is required"})
		return
	}

	payload, err := decodeToken(data["token"])
	if err != nil {
		JSONError(w, err)
		return
	}

	if _, exists := payload["ref"]; !exists {
		fmt.Println(payload)
		JSONError(w, InvalidToken)
		return
	}

	username := payload["username"]

	switch username.(type) {
	case string:
		access, err := newAccessToken(username.(string))
		if err != nil {
			JSONError(w, err)
			return
		}
		Response(w, 200, dict{"access": access})
	default:
		JSONError(w, InvalidToken)
	}
}
