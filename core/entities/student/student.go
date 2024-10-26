package student

import (
	"github.com/joaofilippe/edu-uni-srv/core/entities/class"
	"github.com/joaofilippe/edu-uni-srv/core/entities/grade"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
)

type Student struct {
	id           int
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
