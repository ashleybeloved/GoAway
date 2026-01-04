package middleware

type MiddlewareError string

func (e MiddlewareError) Error() string { return string(e) }

var (
	ErrLoggedOut = MiddlewareError("you are logged out")
)

func CheckAuth(token string) (uint, error) {
	userID, err := ValidateSession(token)
	if err != nil {
		return 0, ErrLoggedOut
	}

	return userID, nil
}
