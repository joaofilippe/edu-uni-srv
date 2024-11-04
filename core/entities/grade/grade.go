package gradeentities

import "github.com/shopspring/decimal"

type Grade struct {
	id        int
	studentId int
	teacherId int
	classId   int
	grade     decimal.Decimal
}

func NewGrade(
	id,
	studentId,
	teacherId,
	classId int,
	grade decimal.Decimal,
) *Grade {
	return &Grade{
		id,
		studentId,
		teacherId,
		classId,
		grade,
	}
}
