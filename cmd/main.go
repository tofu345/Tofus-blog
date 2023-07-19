package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tofu345/Tofus-blog/src"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// ? Check Permissions Middleware
// ? Change post author identification
// ? Format post on post detail page
// ! Show comments
// ? Generate HTML for post_detail page with js like post_list page
// ??? Store post list in localstorage and refresh periodically
// ? Change background image for error page
// ? User Settings Page (not too important rn)
// ! User logged in middleware
// ! Pagination
// ? sort post list by updated_at
// ! Activity Stream
// ! Admin Page
// ! Have post likes be positioned with grid not padding
// ? Expand post on post list page instead of redirect
// ? use varaibles in css
// ! implement something similar to messages in django
// ! Prevent some fields from being changed via update api e.g. likes and views
// ? Upload files
// ? Js rich-text editor e.g. TinyMCE
// ? Convert \n to <br> when displaying posts
// ? Create post list ui with js on frontend

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

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
