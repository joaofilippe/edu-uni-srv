package guardianrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	guardianentities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
)

type GuardianDBModel struct {
	ID        uuid.UUID    `db:"id"`
	UserID    uuid.UUID    `db:"user_id"`
	StudentID uuid.UUID    `db:"student_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Active    bool         `db:"active"`
}

func (g *GuardianDBModel) toEntity() *guardianentities.Guardian {
	return guardianentities.NewGuardian(
		g.ID,
		g.UserID,
		g.StudentID,
		g.CreatedAt,
		g.UpdatedAt.Time,
		g.Active,
	)
}
