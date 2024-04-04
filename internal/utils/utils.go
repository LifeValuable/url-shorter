package utils

import (
	"io"
	"math/rand"
	"net/http"
)

func GetLinkFromBody(request *http.Request) (string, error) {
	buf := make([]byte, 100)
	size, err := request.Body.Read(buf)
	if err != nil {
		if err != io.EOF {
			return "", err
		}
	}
	return string(buf[:size]), nil
}

func GenerateLink() string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	link := make([]rune, 6)
	for i := range link {
		link[i] = chars[rand.Intn(len(chars))]
	}

	return string(link)
}
