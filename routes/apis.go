package routes

import (
	"net/http"
	"strconv"
	"tofs-blog/backend"
	"tofs-blog/posts"
	"tofs-blog/user"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func postListApi(w http.ResponseWriter, r *http.Request) {
	objects, err := posts.GetAll()
	if err != nil {
		JSONResponse(w, 103, err, "Error fetching records")
		return
	}

	JSONResponse(w, 100, objects, "Post List")
}

func postDetailApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		JSONResponse(w, 103, err, "Invalid URL")
		return
	}

	post, e := posts.GetById(id)
	if e != nil {
		JSONResponse(w, 103, err, "Post Not Found")
		return
	}

	JSONResponse(w, 100, post, "Post Detail")
}

func createPostApi(w http.ResponseWriter, r *http.Request) {
	var post backend.Post
	err := JSONDecode(r, &post)
	if err != nil {
		JSONResponse(w, 103, err, "Invalid URL")
	}

	post.Slug = slug.Make(post.Title)

	err_map := posts.Errors(post)
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, "Data Invalid")
		return
	}

	err = posts.Create(&post)
	if err != nil {
		JSONResponse(w, 103, err, "Error Creating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func deletePostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		JSONResponse(w, 103, err, "Invalid URL")
		return
	}

	posts.Delete(id)
	JSONResponse(w, 100, nil, "Success")
}

func updatePostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		JSONResponse(w, 103, err, "Invalid URL")
		return
	}

	var post backend.Post
	post, err = posts.GetById(id)
	if err != nil {
		JSONResponse(w, 103, err, "Post Not Found")
		return
	}

	oldTitle := post.Title
	err = JSONDecode(r, &post)
	if err != nil {
		JSONResponse(w, 103, err, "Invalid POST Data")
		return
	}
	post.ID = id // in case id is passed
	if oldTitle != post.Title {
		post.Slug = slug.Make(post.Title) // Regenerate slug
	}

	errs := posts.Errors(post)
	if len(errs) != 0 {
		JSONResponse(w, 103, errs, "Data Invalid")
		return
	}

	err = posts.Update(&post)
	if err != nil {
		JSONResponse(w, 103, err, "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

// TODO Sign up and sign in api views
func signUpApi(w http.ResponseWriter, r *http.Request) {
	var userObj backend.User
	err := JSONDecode(r, &userObj)
	if err != nil {
		JSONResponse(w, 103, err, "Invalid URL")
		return
	}

	err_map := user.Errors(&userObj)
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, "Data Invalid")
		return
	}

	var pswd string
	pswd, err = backend.HashPassword(userObj.Password)
	if err != nil {
		JSONResponse(w, 103, err, "Error Creating Password Hash")
		return
	}
	userObj.Password = pswd

	if err = user.Create(&userObj); err != nil {
		JSONResponse(w, 103, err, "Error Creating User")
		return
	}

	JSONResponse(w, 100, userObj, "Success")
}

//	{
//	    "author": "tofs",
//	    "message": "Sounds Good"
//	}
func createCommentApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		JSONResponse(w, 103, err, "Invalid URL")
		return
	}

	var post backend.Post
	post, err = posts.GetById(id)
	if err != nil {
		JSONResponse(w, 103, err, "Post Not Found")
		return
	}

	var comment backend.Comment
	err = JSONDecode(r, &comment)
	if err != nil {
		JSONResponse(w, 103, err, "Invalid POST Data")
		return
	}

	post.Comments = append(post.Comments, comment)
	err = posts.Update(&post)
	if err != nil {
		JSONResponse(w, 103, err, "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}
