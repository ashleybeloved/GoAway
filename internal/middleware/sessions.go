package middleware

import (
	"fmt"
	"goaway/internal/repositories"
	"strconv"

	"github.com/google/uuid"
)

func GenerateSession(userID uint) (string, error) {
	sessionToken := uuid.New().String()

	err := repositories.SetSession(sessionToken, userID)
	if err != nil {
		return "", err
	}

	return sessionToken, nil
}

func ValidateSession(token string) (uint, error) {
	idStr, err := repositories.GetSession(token)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error format ID in Redis")
	}

	err = repositories.RefreshSession(token)
	if err != nil {
		return 0, fmt.Errorf("error to refresh session")
	}

	return uint(id), nil
}

func DeleteSession(token string) error {
	err := repositories.DelSession(token)
	if err != nil {
		return err
	}

	return nil
}
