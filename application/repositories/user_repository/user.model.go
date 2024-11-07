package user_repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	userentities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
)

type UserDBModel struct {
	ID        string       `db:"id"`
	Username  string       `db:"username"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	UserType  string       `db:"user_type"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Active    bool         `db:"active"`
}

func (u *UserDBModel) toEntity() (*userentities.User, error) {
	userType, err := enums.ParseUserType(u.UserType)
	if err != nil {
		return &userentities.User{}, err
	}

	return userentities.NewUser(
		uuid.MustParse(u.ID),
		u.Email,
		u.Username,
		u.Password,
		userType,
		nil,
		u.CreatedAt,
		u.UpdatedAt.Time,
		u.Active,
	)
}
