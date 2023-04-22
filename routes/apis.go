package routes

import (
	"net/http"
	"time"
	"tofs-blog/db"

	"github.com/gosimple/slug"
)

func postListApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	postList := []db.Post{}
	err := db.Get(&postList)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error fetching records")
		return
	}

	JSONResponse(w, 100, postList, "Post List")
}

func postDetailApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	post := db.Post{ID: id}
	err = db.Get(&post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Post Not Found")
		return
	}

	JSONResponse(w, 100, post, "Post Detail")
}

func createPostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var post db.Post
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

	err = db.Create(&post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func deletePostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	db.Delete(&db.Post{ID: id})
	JSONResponse(w, 100, nil, "Success")
}

func updatePostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	post := db.Post{ID: id}
	err = db.Get(&post)
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

	err = db.Update(&post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func createCommentApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id, err := getIdFromRequest(r)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	post := db.Post{ID: id}
	err = db.Get(&post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Post Not Found")
		return
	}

	var comment db.Comment
	err = JSONDecode(r, &comment)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidPOSTData)
		return
	}

	post.Comments = append(post.Comments, comment)
	err = db.Update(&post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func userSignupApi(w http.ResponseWriter, r *http.Request) {
	var user db.User
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
	pswd, err = db.HashPassword(user.Password)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating Password Hash")
		return
	}
	user.Password = pswd

	err = db.Create(&user)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating User")
		return
	}

	JSONResponse(w, 100, user, "Success")
}

func userListApi(w http.ResponseWriter, r *http.Request) {
	users := []db.User{}
	err := db.Get(&users)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error fetching records")
		return
	}

	JSONResponse(w, 100, users, "User List")
}

// Check username and password
// Generate token and token expiry date and store
// return token
func userLoginApi(w http.ResponseWriter, r *http.Request) {
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

	user, err := db.GetUserByEmail(userData["email"])
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error")
		return
	}

	if !db.CheckPasswordHash(userData["password"], user.Password) {
		JSONResponse(w, 103, "Incorrect Password", "Error")
		return
	}

	user.AccessToken = db.GenerateToken(user.Email)
	user.TokenExpiryDate = time.Now().Add(7 * 24 * time.Hour) // Expires in 7 days

	err = db.Update(&user)
	if err != nil {
		JSONResponse(w, 103, "Error generating token", "Error")
		return
	}

	JSONResponse(w, 100, map[string]any{"user": user, "token": user.AccessToken}, "Success")
}
