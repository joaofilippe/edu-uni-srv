package userentities

import (
	"strings"

	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	"github.com/joaofilippe/edu-uni-srv/domain/interfaces"
)

type CreateUser struct {
	id          uuid.UUID
	username    string
	password    string
	email       string
	userType    enums.UserType
	userDetails interfaces.IUserDetails
}

func NewCreateUser(
	username string,
	password string,
	email string,
	userType enums.UserType,
	userDetails interfaces.IUserDetails,
) *CreateUser {
	return &CreateUser{
		uuid.UUID{},
		username,
		password,
		email,
		userType,
		userDetails,
	}
}

func (c *CreateUser) Username() string {
	return c.username
}

func (c *CreateUser) Password() string {
	return c.password
}

func (c *CreateUser) SetPassword(password string) {
	c.password = password
}

func (c *CreateUser) Email() string {
	return c.email
}

func (c *CreateUser) UserType() enums.UserType {
	return c.userType
}

func (c *CreateUser) UserDetails() interfaces.IUserDetails {
	return c.userDetails
}

func (c *CreateUser) ID() uuid.UUID {
	return c.id
}

func (c *CreateUser) EmptyID() bool {
	return c.id == uuid.UUID{}
}

func (c *CreateUser) SetID(id uuid.UUID) {
	c.id = id
}

func (c *CreateUser) ValidateEmail() bool {
	if len(c.email) < 3 && len(c.email) > 254 {
		return false
	}

	at := strings.LastIndex(c.email, "@")
	if at < 1 || at+1 >= len(c.email)-1 {
		return false
	}

	domain := c.email[at+1:]
	dot := strings.LastIndex(domain, ".")
	if dot < 1 || dot+1 >= len(domain)-1 {
		return false
	}

	return true
}
