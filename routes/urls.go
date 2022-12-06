package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	url      string
	method   string
	function func(http.ResponseWriter, *http.Request)
}

func RegisterRoutes(r *mux.Router) {
	// Add new routes here
	routes := []Route{
		{url: "/", method: "GET", function: homeView},
		{url: "/posts", method: "GET", function: postListView},
		{url: "/create", method: "POST", function: createPostView},
		{url: "/delete/{id}", method: "DELETE", function: deletePostView},
		{url: "/update/{id}", method: "PUT", function: updatePostView},
	}

	for _, route := range routes {
		r.HandleFunc(route.url, route.function).Methods(route.method)
	}
}
