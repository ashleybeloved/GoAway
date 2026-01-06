package services

import (
	"goaway/internal/models"
	"goaway/internal/repositories"
	"math/rand"
	"strings"
)

type LinkError string

func (e LinkError) Error() string { return string(e) }

var (
	ErrCreateLink     = LinkError("could not create link")
	ErrURLNotExists   = LinkError("page not found")
	ErrInvalidRequest = LinkError("invalid json request")
	ErrNotLink        = LinkError("invalid link")
	ErrLinkNotFound   = LinkError("link not found")
	ErrGetLinks       = LinkError("could not get links")
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

	err := repositories.SetLink(url, string(shortUrl), userID)
	if err != nil {
		return "", ErrCreateLink
	}

	err = repositories.CreateLink(url, string(shortUrl), userID)
	if err != nil {
		return "", ErrCreateLink
	}

	return string(shortUrl), nil
}

func Redirect(shortUrl string) (string, error) {
	var link *models.Link

	url, err := repositories.GetLink(shortUrl)
	if err != nil || url == "" {
		link, err = repositories.GetLinkByShortURL(shortUrl)
		if err != nil || link == nil {
			return "", ErrURLNotExists
		}

		go repositories.AddClick(shortUrl)

		return link.URL, nil
	}

	go repositories.AddClick(shortUrl)

	return url, nil
}

func DelLink(shortUrl string, userID uint) error {
	err := repositories.DelLinkByUser(shortUrl, userID)
	if err != nil {
		return ErrLinkNotFound
	}

	return nil
}

func Link(shortUrl string, userID uint) (*models.Link, error) {
	link, err := repositories.GetLinkByShortURLAndUser(shortUrl, userID)
	if err != nil {
		return nil, ErrLinkNotFound
	}

	return link, nil
}

func Links(userID uint) ([]models.Link, error) {
	links, err := repositories.GetAllUserLinks(userID)
	if err != nil {
		return nil, ErrGetLinks
	}

	return links, nil
}
