package user

import "errors"

type User struct {
	id       int
	username string
	password string
	email    string
}

func NewUser(
	id int,
	username string,
	password string,
	email string,
) (*User, error) {
	user := &User{id,
		username,
		password,
		email,
	}

	if !user.validateEmail() {
		return nil, errors.New("invalid email")
	}

	if !user.validateUsername() {
		return nil, errors.New("invalid username")
	}

	return user, nil
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Email() string {
	return u.email
}

func (u *User) validateUsername() bool {
	return true
}

func (u *User) validateEmail() bool {
	return true
}
