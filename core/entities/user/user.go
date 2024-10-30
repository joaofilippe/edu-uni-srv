package user

import (
	"errors"
	
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/interfaces"
)

type User struct {
	id          int
	username    string
	password    string
	email       string
	userType    enums.UserType
	userDetails interfaces.IUserDetails
}

func NewUser(
	id int,
	username string,
	password string,
	email string,
	userType enums.UserType,
	userDetails interfaces.IUserDetails,
) (*User, error) {
	user := &User{id,
		username,
		password,
		email,
		userType,
		userDetails,
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

func (u *User) UserType() enums.UserType {
	return u.userType
}

func (u *User) UserDetails() interfaces.IUserDetails {
	return u.userDetails
}

func (u *User) validateUsername() bool {
	return true
}

func (u *User) validateEmail() bool {
	return true
}
