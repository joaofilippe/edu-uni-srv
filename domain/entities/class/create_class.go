package classentities

import "github.com/google/uuid"

type CreateClass struct {
	id         uuid.UUID
	name       string
	teacherID  uuid.UUID
	studentsIDs []uuid.UUID
	contentsIDs []uuid.UUID
}

func NewCreateClass(
	name string,
	teacherID uuid.UUID,
	studentsIDs []uuid.UUID,
	contentsIDs []uuid.UUID,
) *CreateClass {
	return &CreateClass{

		name:       name,
		teacherID:  teacherID,
		studentsIDs: studentsIDs,
		contentsIDs: contentsIDs,
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

func (c *CreateClass) StudentsIDs() []uuid.UUID {
	return c.studentsIDs
}

func (c *CreateClass) ContentsIDs() []uuid.UUID {
	return c.contentsIDs
}

func (c *CreateClass) SetID(id uuid.UUID) {
	c.id = id
}
