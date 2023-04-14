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
		{url: "/api/posts", methods: []string{"GET"}, function: postListApi},
		{url: "/api/posts/{id}", methods: []string{"GET"}, function: postDetailApi},
		{url: "/api/post/{id}/comments", methods: []string{"POST"}, function: createCommentApi},
		{url: "/api/create", methods: []string{"POST"}, function: createPostApi},
		{url: "/api/delete/{id}", methods: []string{"DELETE"}, function: deletePostApi},
		{url: "/api/update/{id}", methods: []string{"PUT"}, function: updatePostApi},

		{url: "/signup", methods: []string{"POST"}, function: signUpApi},
		{url: "/posts/{slug}", methods: []string{"GET"}, function: postDetailView},
	}

	// Static Files
	fs := http.FileServer(http.Dir("./staticfiles/"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))

	for _, route := range routes {
		r.HandleFunc(route.url, route.function).Methods(route.methods...)
	}

	// Custom 404 Handler
	r.NotFoundHandler = http.HandlerFunc(NotFound404Handler)
}
