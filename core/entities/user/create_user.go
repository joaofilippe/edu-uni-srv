package user

import (
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/interfaces"
)

type CreateUser struct {
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

func (c *CreateUser) Email() string {
	return c.email
}

func (c *CreateUser) UserType() enums.UserType {
	return c.userType
}

func (c *CreateUser) UserDetails() interfaces.IUserDetails {
	return c.userDetails
}
