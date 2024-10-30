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
