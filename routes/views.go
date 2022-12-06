package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"tofus-blog/backend"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = backend.GetDB()
}

func homeView(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, 200, map[string]string{"message": "Success!"}, "")
}

func postListView(w http.ResponseWriter, r *http.Request) {
	var posts = []backend.Post{}
	query := db.Find(&posts)

	if query.Error != nil {
		JSONResponse(w, 103, query.Error, "Error fetching records")
	}

	JSONResponse(w, 100, posts, "Post List")
}

func createPostView(w http.ResponseWriter, r *http.Request) {
	var post backend.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	err_map, err := backend.CreatePost(post)
	if err != nil {
		JSONResponse(w, 103, err_map, err.Error())
	} else {
		JSONResponse(w, 100, post, "Post Created Successfully")
	}
}

func deletePostView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	backend.DeletePost(id)
	JSONResponse(w, 100, nil, "Post Deleted Successfully")
}

func updatePostView(w http.ResponseWriter, r *http.Request) {
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
