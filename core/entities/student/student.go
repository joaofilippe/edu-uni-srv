package student

import "github.com/joaofilippe/edu-uni-srv/core/enums"

type Student struct {
	id           int
	name         string
	age          int
	classes      []string
	disabilities []enums.Disability
	guardian     string
	address      string
	phone        string
}

func NewStudent(
	id int,
	name string,
	age int,
	classes []string,
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
		disabilities,
		guardian,
		address,
		phone,
	}
}
