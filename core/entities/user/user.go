package userentities

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/interfaces"
)

type User struct {
	id          uuid.UUID
	username    string
	password    string
	email       string
	userType    enums.UserType
	userDetails interfaces.IUserDetails
	createdAt   time.Time
	updatedAt   time.Time
	active      bool
}

func NewUser(
	id uuid.UUID,
	username string,
	password string,
	email string,
	userType enums.UserType,
	userDetails interfaces.IUserDetails,
	createdAt time.Time,
	updatedAt time.Time,
	active bool,
) (*User, error) {
	user := &User{id,
		username,
		password,
		email,
		userType,
		userDetails,
		createdAt,
		updatedAt,
		active,
	}

	if !user.validateEmail() {
		return nil, errors.New("invalid email")
	}

	if !user.validateEmail() {
		return nil, errors.New("invalid username")
	}

	return user, nil
}

func (u *User) ID() uuid.UUID {
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

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) Active() bool {
	return u.active
}

func (u *User) CopyWith(
	username *string,
	password *string,
	email *string,
	userType *enums.UserType,
	userDetails *interfaces.IUserDetails,
) *User {
	newUser := *u

	if username != nil {
		newUser.username = *username
	}

	if password != nil {
		newUser.password = *password
	}

	if email != nil {
		newUser.email = *email
	}

	if userType != nil {
		newUser.userType = *userType
	}

	if userDetails != nil {
		newUser.userDetails = *userDetails
	}

	return &newUser
}

func (u *User) SetUserDetails(userDetails interfaces.IUserDetails) {
	u.userDetails = userDetails
}

func (u *User) validateEmail() bool {
	if len(u.email) < 3 && len(u.email) > 254 {
		return false
	}

	at := strings.LastIndex(u.email, "@")
	if at < 1 || at+1 >= len(u.email)-1 {
		return false
	}

	domain := u.email[at+1:]
	dot := strings.LastIndex(domain, ".")
	if dot < 1 || dot+1 >= len(domain)-1 {
		return false
	}

	return true
}
