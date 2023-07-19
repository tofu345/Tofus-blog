package src

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	url      string
	methods  []string
	function func(http.ResponseWriter, *http.Request)
}

func RegisterRoutes(r *mux.Router) {
	// Add new routes here
	routes := []Route{
		// Views
		{url: "/", methods: []string{"GET"}, function: homeView},
		{url: "/login", methods: []string{"GET"}, function: loginView},
		{url: "/signup", methods: []string{"GET"}, function: signUpView},
		{url: "/posts/{id}/{slug}", methods: []string{"GET"}, function: postDetailView},

		// Api's
		{url: "/api/posts", methods: []string{"GET"}, function: postList},
		{url: "/api/posts/{id}", methods: []string{"GET"}, function: postDetail},
		{url: "/api/posts/{id}/comments", methods: []string{"POST"}, function: createComment},
		{url: "/api/create", methods: []string{"POST"}, function: createPost},
		{url: "/api/delete/{id}", methods: []string{"DELETE"}, function: deletePost},
		{url: "/api/update/{id}", methods: []string{"PUT"}, function: updatePost},

		{url: "/api/users", methods: []string{"GET"}, function: userList},
		{url: "/api/signup", methods: []string{"POST"}, function: userSignup},
		{url: "/api/token", methods: []string{"POST"}, function: getAccessToken},
		{url: "/api/token/refresh", methods: []string{"POST"}, function: refreshAccessToken},
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
