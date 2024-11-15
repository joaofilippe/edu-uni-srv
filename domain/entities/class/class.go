package classentities

import "github.com/google/uuid"

type Class struct {
	id         uuid.UUID
	name       string
	teacherID  uuid.UUID
	studentsID []uuid.UUID
	contentsID []uuid.UUID
}

func NewClass(
	id int,
	name string,
	teacherId int,
	studentsId []int,
	schedule []string,
) *Class {
	return &Class{
		id,
		name,
		teacherId,
		studentsId,
		schedule,
	}
}

func (c *Class) ID() int {
	return c.id
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) TeacherID() int {
	return c.teacherID
}

func (c *Class) StudentsID() []int {
	return c.studentsID
}

func (c *Class) Schedule() []string {
	return c.schedule
}
