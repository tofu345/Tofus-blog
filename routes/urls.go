package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	url      string
	methods  []string
	function func(http.ResponseWriter, *http.Request)
}

func Register(r *mux.Router) {
	// Add new routes here
	routes := []Route{
		// Views
		{url: "/", methods: []string{"GET"}, function: homeView},

		// Api's
		{url: "/posts", methods: []string{"GET"}, function: postListApi},
		{url: "/posts/{id}", methods: []string{"GET"}, function: postDetailApi},
		{url: "/create", methods: []string{"POST"}, function: createPostApi},
		{url: "/delete/{id}", methods: []string{"DELETE"}, function: deletePostApi},
		{url: "/update/{id}", methods: []string{"PUT"}, function: updatePostApi},
	}

	// Static Files
	fs := http.FileServer(http.Dir("./staticfiles/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))

	for _, route := range routes {
		r.HandleFunc(route.url, route.function).Methods(route.methods...)
	}
}
