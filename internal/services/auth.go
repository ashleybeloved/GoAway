package services

import "unicode"

type AuthError string

func (e AuthError) Error() string { return string(e) }

var (
	ErrInvalidChars     = AuthError("invalid chars in login")
	ErrLoginTooShort    = AuthError("login too short")
	ErrPasswordTooShort = AuthError("password too short")
)

func Reg(login string, password string) error {
	// Check for invalid chars in login

	for _, r := range login {
		if !unicode.Is(unicode.Latin, r) || !unicode.IsDigit(r) {
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

	return nil
}
