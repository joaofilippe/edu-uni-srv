package student

import (
	"github.com/joaofilippe/edu-uni-srv/core/entities/class"
	"github.com/joaofilippe/edu-uni-srv/core/entities/grade"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
)

type Student struct {
	id           int
	userId int
	name         string
	age          int
	classes      []class.Class
	grades       []grade.Grade
	disabilities []enums.Disability
	guardian     string
	address      string
	phone        string
}

func NewStudent(
	id int,
	userId int,
	name string,
	age int,
	classes []class.Class,
	grades []grade.Grade,
	disabilities []enums.Disability,
	guardian string,
	address string,
	phone string,
) *Student {
	return &Student{
		id,
		userId,
		name,
		age,
		classes,
		grades,
		disabilities,
		guardian,
		address,
		phone,
	}
}

func (s *Student) Id() int {
	return s.id
}

func (s *Student) UserId() int {
	return s.userId
}

func (s *Student) Name() string {
	return s.name
}

func (s *Student) Age() int {
	return s.age
}

func (s *Student) Classes() []class.Class {
	return s.classes
}

func (s *Student) Grades() []grade.Grade {
	return s.grades
}

func (s *Student) Disabilities() []enums.Disability {
	return s.disabilities
}

func (s *Student) Guardian() string {
	return s.guardian
}

func (s *Student) Address() string {
	return s.address
}

func (s *Student) Phone() string {
	return s.phone
}