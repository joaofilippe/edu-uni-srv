package classentities

type Class struct {
	id         int
	name       string
	teacherId  int
	studentsId []int
	schedule   []string
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
	return c.teacherId
}

func (c *Class) StudentsID() []int {
	return c.studentsId
}

func (c *Class) Schedule() []string {
	return c.schedule
}
