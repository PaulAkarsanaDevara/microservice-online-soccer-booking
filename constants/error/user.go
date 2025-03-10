package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordInCorrect    = errors.New("password Incorrect")
	ErrUsernameExists       = errors.New("username Already Exists")
	ErrEmailExists          = errors.New("email Already Exists")
	ErrPasswordDoesNotMatch = errors.New("password Does Not Match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameExists,
	ErrPasswordDoesNotMatch,
}
