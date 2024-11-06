package errors

import "errors"

var (
	ErrUserTypeNotFound       = errors.New("user type not found")
	ErrUserIDNot              = errors.New("user id not provided")
	ErrUserInvalidEmail       = errors.New("invalid email")
	ErrUserNoEmail            = errors.New("no email provided")
	ErrUserInvalidUsername    = errors.New("invalid username")
	ErrUserNoUsername         = errors.New("no username provided")
	ErrUserInvalidPassword    = errors.New("invalid password")
	ErrUserNoPassword         = errors.New("no password provided")
	ErrUserInvalidUserType    = errors.New("invalid user type")
	ErrUserNoUserType         = errors.New("no user type provided")
	ErrUserInvalidUserDetails = errors.New("invalid user details")
	ErrUserNoUserDetails      = errors.New("no user details provided")
	ErrUserUserNotFound       = errors.New("user not found")
	ErrUserUserAlreadyExists  = errors.New("user already exists")
)

var (
	ErrAdminIDNot     = errors.New("admin id not provided")
	ErrAdminUserIDNot = errors.New("user id not provided")
)
