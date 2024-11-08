package userrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	userentities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
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

func fromUserEntity(entity *userentities.User) *UserDBModel {
	return &UserDBModel{
		ID:        entity.ID().String(),
		Username:  entity.Username(),
		Email:     entity.Email(),
		Password:  entity.Password(),
		UserType:  entity.UserType().String(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: sql.NullTime{Time: entity.UpdatedAt(), Valid: true},
		Active:    entity.Active(),
	}
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

type CreateUserDBModel struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	UserType string    `db:"user_type"`
}

func fromCreateEntity(entity *userentities.CreateUser) *CreateUserDBModel {
	return &CreateUserDBModel{
		ID:       entity.ID(),
		Username: entity.Username(),
		Email:    entity.Email(),
		Password: entity.Password(),
		UserType: entity.UserType().String(),
	}
}
