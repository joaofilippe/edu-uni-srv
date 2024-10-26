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

	if !user.ValidateEmail() {
		return nil, errors.New("invalid email")
	}

	if !user.ValidateUsername() {
		return nil, errors.New("invalid username")
	}

	return user, nil
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) ValidateUsername() bool {
	return true
}

func (u *User) ValidateEmail() bool {
	return true
}
