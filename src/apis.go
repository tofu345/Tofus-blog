package src

import (
	"net/http"
	"time"

	"github.com/gosimple/slug"
)

func getPostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	postList := []Post{}
	err := db.First(&postList).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error fetching records")
		return
	}

	JSONResponse(w, 100, postList, "Post List")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Post Not Found")
		return
	}

	JSONResponse(w, 100, post, "Post Detail")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var post Post
	err := JSONDecode(r, &post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
	}

	post.Slug = slug.Make(post.Title)
	err_map := post.Errors()
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, InvalidData)
		return
	}

	err = db.Create(&post).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	db.Delete(&Post{ID: id})
	JSONResponse(w, 100, nil, "Success")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Post Not Found")
		return
	}

	oldTitle := post.Title
	err = JSONDecode(r, &post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidPOSTData)
		return
	}

	post.ID = id // in case id is passed
	if oldTitle != post.Title {
		post.Slug = slug.Make(post.Title) // Regenerate slug
	}

	errs := post.Errors()
	if len(errs) != 0 {
		JSONResponse(w, 103, errs, InvalidData)
		return
	}

	err = db.Save(&post).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	post := Post{ID: id}
	err = db.First(&post).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Post Not Found")
		return
	}

	var comment Comment
	err = JSONDecode(r, &comment)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidPOSTData)
		return
	}

	post.Comments = append(post.Comments, comment)
	err = db.Save(&post).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func userSignup(w http.ResponseWriter, r *http.Request) {
	var user User
	err := JSONDecode(r, &user)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidPOSTData)
		return
	}

	err_map := user.Errors()
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, InvalidData)
		return
	}

	var pswd string
	pswd, err = HashPassword(user.Password)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating Password Hash")
		return
	}
	user.Password = pswd

	err = db.Create(&user).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating User")
		return
	}

	JSONResponse(w, 100, user, "Success")
}

func userList(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	err := db.First(&users).Error
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error fetching records")
		return
	}

	JSONResponse(w, 100, users, "User List")
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userData := make(map[string]string)
	err := JSONDecode(r, &userData)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidPOSTData)
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
		JSONResponse(w, 103, errors, InvalidData)
		return
	}

	var user User
	err = db.First(&user, "email = ?", userData["email"]).Error
	if err != nil {
		if err.Error() == RecordNotFound {
			JSONResponse(w, 103, LoginError, "Error")
		} else {
			JSONResponse(w, 103, err.Error(), "Error")
		}
		return
	}

	if !CheckPasswordHash(userData["password"], user.Password) {
		JSONResponse(w, 103, LoginError, "Error")
		return
	}

	user.AccessToken = GenerateToken(user.Email)
	user.TokenExpiryDate = time.Now().Add(7 * 24 * time.Hour) // Token expires in 7 days

	err = db.Save(&user).Error
	if err != nil {
		JSONResponse(w, 103, "Error generating token", "Error")
		return
	}

	JSONResponse(w, 100, map[string]any{"token": user.AccessToken}, "Success")
}
