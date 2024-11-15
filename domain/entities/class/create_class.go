package classentities

import "github.com/google/uuid"

type CreateClass struct {
	id         uuid.UUID
	name       string
	teacherID  uuid.UUID
	studentsID []uuid.UUID
	contentsID []uuid.UUID
}

func NewCreateClass(
	name string,
	teacherID uuid.UUID,
	studentsID []uuid.UUID,
	contentsID []uuid.UUID,
) *CreateClass {
	return &CreateClass{

		name:       name,
		teacherID:  teacherID,
		studentsID: studentsID,
		contentsID: contentsID,
	}
}

func (c *CreateClass) ID() uuid.UUID {
	return c.id
}

func (c *CreateClass) Name() string {
	return c.name
}

func (c *CreateClass) TeacherID() uuid.UUID {
	return c.teacherID
}

func (c *CreateClass) StudentsID() []uuid.UUID {
	return c.studentsID
}

func (c *CreateClass) ContentsID() []uuid.UUID {
	return c.contentsID
}

func (c *CreateClass) SetID(id uuid.UUID) {
	c.id = id
}
