package teacherrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	"github.com/lib/pq"
)

type TeacherDBModel struct {
	ID         uuid.UUID      `db:"id"`
	UserID     uuid.UUID      `db:"user_id"`
	ClassesIDs pq.StringArray `db:"classes_ids"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at"`
	Active     bool           `db:"active"`
}

func (t *TeacherDBModel) toEntity() *teacherentities.Teacher {
	classesIDs := []uuid.UUID{}
	for _, classID := range t.ClassesIDs {
		classesIDs = append(classesIDs, uuid.MustParse(classID))
	}

	return teacherentities.NewTeacher(
		t.ID,
		t.UserID,
		classesIDs,
		t.CreatedAt,
		t.UpdatedAt.Time,
		t.Active,
	)
}

type CreateTeacherDbModel struct {
	ID         uuid.UUID      `db:"id"`
	UserID     uuid.UUID      `db:"user_id"`
	ClassesIDs pq.StringArray `db:"classes_ids"`	
}

func (t *CreateTeacherDbModel) toEntity() *teacherentities.Teacher {
	classesIDs := []uuid.UUID{}
	for _, classID := range t.ClassesIDs {
		classesIDs = append(classesIDs, uuid.MustParse(classID))
	}

	return teacherentities.NewTeacher(
		t.ID,
		t.UserID,
		classesIDs,
		time.Time{},
		time.Time{},
		true,
	)
}

func (t *CreateTeacherDbModel) fromEntity(teacher *teacherentities.CreateTeacher) {
	classesIDs := []string{}
	for _, classID := range teacher.ClassesIDs() {
		classesIDs = append(classesIDs, classID.String())
	}

	t.ID = teacher.ID()
	t.UserID = teacher.UserID()
	t.ClassesIDs = classesIDs
}
