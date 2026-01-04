package services

import (
	"goaway/internal/middleware"
	"goaway/internal/repositories"
	"goaway/pkg"
	"unicode"
)

type AuthError string

func (e AuthError) Error() string { return string(e) }

var (
	ErrInvalidChars     = AuthError("invalid chars in login")
	ErrLoginTooShort    = AuthError("login too short")
	ErrPasswordTooShort = AuthError("password too short")
	ErrUserExists       = AuthError("user with this login already exists")
	ErrUserNotExists    = AuthError("user with this login not exists")
	ErrHashPassword     = AuthError("password was not hashed")
	ErrCreateUser       = AuthError("user was not created")
	ErrInvalidPassword  = AuthError("invalid password")
	ErrGenerateSession  = AuthError("session was not created")
	ErrDelSession       = AuthError("session was not deleted")
)

func Reg(login string, password string) error {
	// Check for invalid chars in login

	for _, r := range login {
		if !unicode.Is(unicode.Latin, r) && !unicode.IsDigit(r) {
			return ErrInvalidChars
		}
	}

	// Check for lenght login & password

	if len(login) < 3 {
		return ErrLoginTooShort
	}

	if len(password) < 6 {
		return ErrPasswordTooShort
	}

	// Check for avaiblity login in DB

	user, err := repositories.FindUserByLogin(login)
	if err == nil && user != nil {
		return ErrUserExists
	}

	// Hash password

	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		return ErrHashPassword
	}

	// Add user in DB

	err = repositories.CreateUser(login, hashedPassword)
	if err != nil {
		return ErrCreateUser
	}

	return nil
}

func Login(login string, password string) (string, error) {
	user, err := repositories.FindUserByLogin(login)
	if err != nil || user == nil {
		return "", ErrUserNotExists
	}

	err = pkg.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return "", ErrInvalidPassword
	}

	sessionToken, err := middleware.GenerateSession(user.ID)
	if err != nil {
		return "", ErrGenerateSession
	}

	return sessionToken, nil
}

func Logout(token string) error {
	err := middleware.DeleteSession(token)
	if err != nil {
		return ErrDelSession
	}

	return nil
}
