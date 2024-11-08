package studentrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	studententities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
)

type StudentDBModel struct {
	ID           uuid.UUID    `db:"id"`
	UserID       uuid.UUID    `db:"user_id"`
	Age          int          `db:"age"`
	Classes      []uuid.UUID  `db:"classes"`
	Grades       []uuid.UUID  `db:"grades"`
	Disabilities []string  `db:"disabilities"`
	GuardianID   uuid.UUID       `db:"guardian_id"`
	Address      string       `db:"address"`
	Phone        string       `db:"phone"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	Active       bool         `db:"active"`
}

func (s *StudentDBModel) toEntity() *studententities.Student {
	disabilites := []enums.Disability{}

	for _, d := range s.Disabilities {
		disabilites = append(disabilites, enums.ParseDisability(d))
	}

	return studententities.NewStudent(
		s.ID,
		s.UserID,
		s.Age,
		s.Classes,
		s.Grades,
		disabilites,
		s.GuardianID,
		s.Address,
		s.Phone,
		s.CreatedAt,
		s.UpdatedAt.Time,
		s.Active,
	)

}
