package main

import (
	"log"
	"net/http"
	"os"
	"tofus-blog/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := "8005"
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	routes.RegisterApi(r)

	log.Print("Listening on port ", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}
