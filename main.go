package main

import (
	"log"
	"net/http"
	"os"
	"tofs-blog/src"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ! Script to give admin user permissions
// ! Activity Stream
// ! Admin Page
// ! Have post likes be positioned with grid not padding
// ! Create Post UI
// ? Expand post on post list page instead of redirect
// ? use varaibles in css
// ! Permissions to view user list api because of passwords
// ! implement something similar to messages in django
// ! Prevent some fields from being changed via update api e.g. likes and views
// ? Upload files
// ? Js rich-text editor e.g. TinyMCE
// ? Convert \n to <br> when displaying posts
// ? Create post list ui with js on frontend

func main() {
	port := "8005"
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	src.RegisterRoutes(r)

	log.Print("Listening on port http://localhost:", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
