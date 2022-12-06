package main

import (
	"log"
	"net/http"
	"tofus-blog/routes"

	"github.com/gorilla/mux"
)

func main() {
	port := "8005"
	r := mux.NewRouter()

	routes.RegisterRoutes(r)

	log.Print("Listening on port ", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}
