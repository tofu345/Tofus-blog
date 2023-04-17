package routes

import (
	"fmt"
	"net/http"
	"tofs-blog/db"

	"github.com/gosimple/slug"
)

func postListApi(w http.ResponseWriter, r *http.Request) {
	postList := []db.Post{}
	err := db.Get(&postList)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error fetching records")
		return
	}

	JSONResponse(w, 100, postList, "Post List")
}

func postDetailApi(w http.ResponseWriter, r *http.Request) {
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

	err = db.Create(post)
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

	err = db.Update(post)
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
	err = db.Update(post)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Updating Post")
		return
	}

	JSONResponse(w, 100, post, "Success")
}

func votePostApi(w http.ResponseWriter, r *http.Request) {
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

	data := make(map[string]string)
	err = JSONDecode(r, &data)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidPOSTData)
		return
	}

	if _, exists := data["vote"]; !exists {
		JSONResponse(w, 103, "Vote option must be present", InvalidPOSTData)
		return
	}

	if data["vote"] == "like" {
		post.Likes++
	} else if data["vote"] == "dislike" {
		post.Likes--
	} else {
		JSONResponse(w, 103, fmt.Sprintf("%v is not a valid option", data["vote"]), InvalidPOSTData)
		return
	}

	db.Update(post)
	JSONResponse(w, 100, post, "Vote Successful")
}

// TODO Sign up and sign in api views
func signUpApi(w http.ResponseWriter, r *http.Request) {
	var userObj db.User
	err := JSONDecode(r, &userObj)
	if err != nil {
		JSONResponse(w, 103, err.Error(), InvalidURL)
		return
	}

	err_map := userObj.Errors()
	if len(err_map) != 0 {
		JSONResponse(w, 103, err_map, InvalidData)
		return
	}

	var pswd string
	pswd, err = db.HashPassword(userObj.Password)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating Password Hash")
		return
	}
	userObj.Password = pswd

	err = db.Create(userObj)
	if err != nil {
		JSONResponse(w, 103, err.Error(), "Error Creating User")
		return
	}

	JSONResponse(w, 100, userObj, "Success")
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
