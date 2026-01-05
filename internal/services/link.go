package services

import (
	"goaway/internal/repositories"
	"math/rand"
	"strings"
)

type LinkError string

func (e LinkError) Error() string { return string(e) }

var (
	ErrCreateLink     = LinkError("could not create link")
	ErrURLNotExists   = LinkError("url not exists")
	ErrInvalidRequest = LinkError("invalid json request")
	ErrNotLink        = LinkError("invalid link")
)

func New(url string, userID uint) (string, error) {
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

	err := repositories.CreateLink(url, string(shortUrl), userID)
	if err != nil {
		return "", ErrCreateLink
	}

	return string(shortUrl), nil
}

func Redirect(shortUrl string) (string, error) {
	link, err := repositories.FindURLByShortURL(shortUrl)
	if err != nil || link == nil {
		return "", ErrURLNotExists
	}

	go repositories.AddClick(shortUrl)

	return link.URL, nil
}
