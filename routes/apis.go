package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tofs-blog/backend"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func postListApi(w http.ResponseWriter, r *http.Request) {
	var posts = []backend.Post{}
	query := db.Find(&posts)

	if query.Error != nil {
		JSONResponse(w, 103, query.Error, "Error fetching records")
	}

	JSONResponse(w, 100, posts, "Post List")
}

func postDetailApi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	post, e := backend.GetPostById(id)
	if e != nil {
		JSONResponse(w, 103, e, "Post Not Found")
		return
	}

	JSONResponse(w, 100, post, "Post Detail")
}

func createPostApi(w http.ResponseWriter, r *http.Request) {
	var post backend.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.Slug = slug.Make(post.Title)

	err_map := post.Errors()
	if len(err_map) == 0 {
		if result := db.Create(&post); result.Error != nil {
			JSONResponse(w, 103, result.Error.Error(), "Error Creating Post")
		} else {
			JSONResponse(w, 100, post, "Success")
		}
	} else {
		JSONResponse(w, 103, err_map, "Data Invalid")
	}
}

func deletePostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	backend.DeletePost(id)
	JSONResponse(w, 100, nil, "Success")
}

func updatePostApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Check if post exists
	post, e := backend.GetPostById(id)
	if e != nil {
		JSONResponse(w, 103, e, "Post Not Found")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = id // in case id is passed
	errs := post.Errors()

	// Regenerate slug
	post.Slug = slug.Make(post.Title)

	if len(errs) == 0 {
		if res := db.Save(&post); res.Error != nil {
			JSONResponse(w, 103, res.Error.Error(), "Error")
		} else {
			JSONResponse(w, 100, post, "Success")
		}
	} else {
		JSONResponse(w, 103, errs, "Error")
	}
}

// TODO Sign up and sign in api views
func signUpApi(w http.ResponseWriter, r *http.Request) {
	var user backend.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err_map := user.Errors()
	if len(err_map) == 0 {
		// Hash Password
		if password, err := backend.HashPassword(user.Password); err != nil {
			JSONResponse(w, 103, err, "Error")
			return
		} else {
			user.Password = password
		}

		if result := db.Create(&user); result.Error != nil {
			JSONResponse(w, 103, result.Error.Error(), "Error Creating User")
		} else {
			JSONResponse(w, 100, user, "Success")
		}
	} else {
		JSONResponse(w, 103, err_map, "Data Invalid")
	}
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
		log.Fatal(err)
	}

	// Check if post exists
	post, e := backend.GetPostById(id)
	if e != nil {
		JSONResponse(w, 103, e, "Post Not Found")
		return
	}

	var comment backend.Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)

	post.Comments = append(post.Comments, comment)
	if res := db.Save(&post); res.Error != nil {
		JSONResponse(w, 103, res.Error.Error(), "Error")
	} else {
		JSONResponse(w, 100, post, "Success")
	}
}
