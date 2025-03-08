package error

import "errors"

var (
	ErrUserNotFound         = errors.New("User not found")
	ErrPasswordInCorrect    = errors.New("Password Incorrect")
	ErrUsernameExists       = errors.New("Username Already Exists")
	ErrPasswordDoesNotMatch = errors.New("Password Does Not Match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameExists,
	ErrPasswordDoesNotMatch,
}
