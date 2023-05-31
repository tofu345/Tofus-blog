package src

import (
	"net/http"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

func getPostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	postList := []Post{}
	err := db.Find(&postList).Error
	if err != nil && err.Error() != RecordNotFound {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, postList, "Post List")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, err := getIdFromRequest(r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, post, "Post Detail")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	user, err := getUserFromRequestApi(w, r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	var post Post
	err = JSONDecode(r, &post)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post.Author = user.ID
	maxChars := 60
	post.Slug = post.Title
	if maxChars < len(post.Slug) {
		post.Slug = post.Slug[:strings.LastIndex(post.Slug[:maxChars], " ")]
	}

	post.Slug = slug.Make(post.Slug)
	err_map := post.Errors()
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, "Error")
		return
	}

	err = db.Create(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	user, err := getUserFromRequestApi(w, r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	if !userHasPerm(user, "delete_post") && post.Author != user.ID {
		JSONError(w, "You do not have permission to delete this")
		return
	}

	err = db.Delete(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, nil, "Success")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	user, err := getUserFromRequestApi(w, r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	id, err := getIdFromRequest(r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	if !userHasPerm(user, "delete_post") && post.Author != user.ID {
		JSONError(w, "You do not have permission to edit this")
		return
	}

	oldTitle := post.Title
	err = JSONDecode(r, &post)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post.ID = id // in case id is passed
	if oldTitle != post.Title {
		post.Slug = slug.Make(post.Title) // Regenerate slug
	}

	errs := post.Errors()
	if len(errs) != 0 {
		JSONResponse(w, 103, errs, "Error")
		return
	}

	err = db.Save(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	var comment Comment
	err = JSONDecode(r, &comment)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	post.Comments = append(post.Comments, comment)
	err = db.Save(&post).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func userSignup(w http.ResponseWriter, r *http.Request) {
	var user User
	err := JSONDecode(r, &user)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	err_map := user.Errors()
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, "Error")
		return
	}

	var pswd string
	pswd, err = HashPassword(user.Password)
	if err != nil {
		JSONError(w, err.Error())
		return
	}
	user.Password = pswd

	err = db.Create(&user).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, user, "Success")
}

func userList(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	err := db.Find(&users).Error
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	JSONResponse(w, 100, users, "User List")
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userData := make(map[string]string)
	err := JSONDecode(r, &userData)
	if err != nil {
		JSONError(w, err.Error())
		return
	}

	errors := map[string]string{}
	if _, exists := userData["email"]; exists == false {
		errors["email"] = "This field is required"
	}
	if _, exists := userData["password"]; exists == false {
		errors["password"] = "This field is required"
	}
	if len(errors) != 0 {
		JSONResponse(w, 103, errors, "Error")
		return
	}

	var user User
	err = db.First(&user, "email = ?", userData["email"]).Error
	if err != nil {
		err_str := err.Error()
		if err_str == RecordNotFound {
			err_str = LoginError
		}
		JSONError(w, err.Error())
		return
	}

	if !CheckPasswordHash(userData["password"], user.Password) {
		JSONError(w, LoginError)
		return
	}

	user.AccessToken = GenerateToken(user.Email)
	user.TokenExpiryDate = time.Now().Add(7 * 24 * time.Hour) // Token expires in 7 days

	err = db.Save(&user).Error
	if err != nil {
		JSONError(w, "Error generating token")
		return
	}

	JSONResponse(w, 100, map[string]any{"token": user.AccessToken}, "Success")
}
