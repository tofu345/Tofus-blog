package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tofus-blog/backend"

	"github.com/gorilla/mux"
)

func postListApi(w http.ResponseWriter, r *http.Request) {
	var posts = []backend.Post{}
	query := db.Find(&posts)

	if query.Error != nil {
		JSONResponse(w, 103, query.Error, "Error fetching records")
	}

	JSONResponse(w, 100, posts, "Post List")
}

func createPostApi(w http.ResponseWriter, r *http.Request) {
	var post backend.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	err_map, err := backend.CreatePost(post)
	if err != nil {
		JSONResponse(w, 103, err_map, err.Error())
	} else {
		JSONResponse(w, 100, post, "Post Created Successfully")
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
	JSONResponse(w, 100, nil, "Post Deleted Successfully")
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
		JSONResponse(w, 103, nil, "Post Not Found")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = id // in case id is passed
	errs := post.Errors()

	if len(errs) == 0 {
		db.Save(&post)
		JSONResponse(w, 100, post, "Post Updated Successfully")
	} else {
		JSONResponse(w, 103, errs, "Error Updating Post")
	}
}
