package main

import (
	"log"
	"net/http"
	"os"
	"tofs-blog/src"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

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
