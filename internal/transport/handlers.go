package transport

import (
	"net/http"
	"url-shorter/internal/utils"

	"github.com/gorilla/mux"
)

func shorten(writer http.ResponseWriter, request *http.Request) {
	link, err := utils.GetLinkFromBody(request)

	if err != nil {
		http.Error(writer, "", http.StatusInternalServerError)
	}

	if link == "" {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	shortenedLink, ok := memory.GetShortenedLink(link)
	if !ok {
		shortenedLink = utils.GenerateLink()
		memory.InsertLink(link, shortenedLink)
	}

	shortenedLinkByte := []byte("http://localhost:8080/" + shortenedLink)

	writer.WriteHeader(http.StatusOK)
	writer.Write(shortenedLinkByte)
}

func redirectShortened(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	shortenedLink, ok := vars["shortened_link"]

	if !ok {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	link, ok := memory.GetLink(shortenedLink)

	if !ok {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	linkByte := []byte(link)
	writer.WriteHeader(http.StatusOK)
	writer.Write(linkByte)
}
