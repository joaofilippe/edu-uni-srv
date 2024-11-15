package classrepository

import (
	"database/sql"

	"github.com/google/uuid"
	classentities "github.com/joaofilippe/edu-uni-srv/domain/entities/class"
	"github.com/lib/pq"
)

type ClassDbModel struct {
	ID          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	TeacherID   uuid.UUID      `db:"teacher_id"`
	StudentsIDs pq.StringArray `db:"students_ids"`
	ContentsIDs pq.StringArray `db:"contents_ids"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

func (c *ClassDbModel) fromEntity(entity *classentities.Class) {
	studentsIDs := []string{}
	contentsIDs := []string{}

	for _, s := range entity.StudentsIDs() {
		studentsIDs = append(studentsIDs, s.String())
	}

	for _, c := range entity.ContentsIDs() {
		contentsIDs = append(contentsIDs, c.String())
	}

	c.ID = entity.ID()
	c.Name = entity.Name()
	c.TeacherID = entity.TeacherID()
	c.StudentsIDs = pq.StringArray(studentsIDs)
	c.ContentsIDs = pq.StringArray(contentsIDs)
	c.CreatedAt = sql.NullTime{Time: entity.CreatedAt(), Valid: true}
	c.UpdatedAt = sql.NullTime{Time: entity.UpdatedAt(), Valid: true}
}

func (c *ClassDbModel) toEntity() *classentities.Class {
	studentsIDs := []uuid.UUID{}
	contentsIDs := []uuid.UUID{}

	for _, s := range c.StudentsIDs {
		id, _ := uuid.Parse(s)
		studentsIDs = append(studentsIDs, id)
	}

	for _, c := range c.ContentsIDs {
		id, _ := uuid.Parse(c)
		contentsIDs = append(contentsIDs, id)
	}

	return classentities.NewClass(
		c.ID,
		c.Name,
		c.TeacherID,
		studentsIDs,
		contentsIDs,
		c.CreatedAt.Time,
		c.UpdatedAt.Time,
	)
}

type CreateClassDbModel struct {
	ID         uuid.UUID      `db:"id"`
	Name       string         `db:"name"`
	TeacherID  uuid.UUID      `db:"teacher_id"`
}

func (c *CreateClassDbModel) fromEntity(entity *classentities.CreateClass) {
	c.ID = entity.ID()
	c.Name = entity.Name()
	c.TeacherID = entity.TeacherID()
}
