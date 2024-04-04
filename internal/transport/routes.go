package transport

import (
	"net/http"
	"url-shorter/internal/models"

	"github.com/gorilla/mux"
)

var memory models.Memory

func RegisterRoutes(router *mux.Router, withDB bool) {
	switch withDB {
	case true:
		memory = &models.Postgres{}
	case false:
		memory = &models.Links{}
	}
	memory.Init()
	router.HandleFunc("/", shorten).Methods(http.MethodPost)
	router.HandleFunc("/{shortened_link}", redirectShortened).Methods(http.MethodGet)
}
