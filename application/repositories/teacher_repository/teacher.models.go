package teacherrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
)

type TeacherDBModel struct {
	ID         uuid.UUID    `json:"id"`
	UserID     uuid.UUID    `json:"user_id"`
	ClassesIDs []uuid.UUID  `json:"classes_ids"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	Active     bool         `json:"active"`
}

func (t *TeacherDBModel) toEntity() *teacherentities.Teacher {
	return teacherentities.NewTeacher(
		t.ID,
		t.UserID,
		t.ClassesIDs,
		t.CreatedAt,
		t.UpdatedAt.Time,
		t.Active,
	)
}
