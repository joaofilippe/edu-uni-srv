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
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Active    bool         `db:"active"`
}

func FromEntity(entity *adminentities.Admin) *AdminDBModel {
	return &AdminDBModel{
		ID:        entity.ID(),
		UserID:    entity.UserID(),
		CreatedAt: entity.CreatedAt(),
		Active:    entity.Active(),
	}
}

func (a *AdminDBModel) toEntity() *adminentities.Admin {
	return adminentities.NewAdmin(
		a.ID,
		a.UserID,
		a.CreatedAt,
		a.UpdatedAt.Time,
		a.Active,
	)
}