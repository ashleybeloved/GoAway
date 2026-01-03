package services

import (
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

	user, err := repositories.FindUserbyLogin(login)
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

func Login(login string, password string) error {
	user, err := repositories.FindUserbyLogin(login)
	if err != nil && user == nil {
		return ErrUserNotExists
	}

	err = pkg.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
