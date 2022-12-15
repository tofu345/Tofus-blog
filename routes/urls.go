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

func RegisterApi(r *mux.Router) {
	// Add new routes here
	routes := []Route{
		// Views
		{url: "/", method: "GET", function: homeView},

		// Api's
		{url: "/posts", method: "GET", function: postListApi},
		{url: "/create", method: "POST", function: createPostApi},
		{url: "/delete/{id}", method: "DELETE", function: deletePostApi},
		{url: "/update/{id}", method: "PUT", function: updatePostApi},
	}

	for _, route := range routes {
		r.HandleFunc(route.url, route.function).Methods(route.method)
	}
}
