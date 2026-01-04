package services

import (
	"goaway/internal/middleware"
	"goaway/internal/repositories"
	"log"
	"math/rand"
	"strings"
)

type LinkError string

func (e LinkError) Error() string { return string(e) }

var (
	ErrValidateSession = LinkError("invalid session")
	ErrCreateLink      = LinkError("could not create link")
	ErrURLNotExists    = LinkError("url not exists")
	ErrInvalidRequest  = LinkError("invalid json request")
	ErrNotLink         = LinkError("invalid link (prefix http:// or https:// not found)")
)

func New(url string, token string) (string, error) {
	userid, err := middleware.ValidateSession(token)
	if err != nil {
		return "", ErrValidateSession
	}

	if len(url) == 0 {
		return "", ErrInvalidRequest
	}

	if !strings.Contains(url, ".") {
		return "", ErrNotLink
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortUrl := make([]byte, 6)
	for i := range shortUrl {
		shortUrl[i] = charset[rand.Intn(len(charset))]
	}

	err = repositories.CreateLink(url, string(shortUrl), userid)
	if err != nil {
		return "", ErrCreateLink
	}

	return string(shortUrl), nil
}

func Redirect(shortUrl string) (string, error) {
	link, err := repositories.FindURLByShortURL(shortUrl)
	log.Println(link, err)
	if err != nil || link == nil {
		return "", ErrURLNotExists
	}

	return link.URL, nil
}
