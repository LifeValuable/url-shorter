package app

import (
	"log"
	"net/http"
	"url-shorter/internal/transport"

	"github.com/gorilla/mux"
)

func Run(withDB bool) {
	router := mux.NewRouter()
	transport.RegisterRoutes(router, withDB)
	http.Handle("/", router)
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
