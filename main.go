package main

import (
	"log"
	"net/http"
	"os"
	"tofs-blog/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ! Have post likes be positioned with grid not padding
// ? Change db error messages

func main() {
	port := "8005"
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	routes.Register(r)

	log.Print("Listening on port http://localhost:", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
