package adminrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type AdminDBModel struct {
	ID        uuid.UUID    `db:"id"`
	UserID    uuid.UUID    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"updated_at"`
	Active    bool      `db:"active"`
}

func(a *AdminDBModel) toEntity() *adminentities.Admin {
	return adminentities.NewAdmin(
		a.ID,
		a.UserID,
		a.CreatedAt,
		a.UpdatedAt.Time,
		a.Active,
	)
}