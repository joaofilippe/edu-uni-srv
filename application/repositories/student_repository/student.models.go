package studentrepository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	studententities "github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	"github.com/lib/pq"
)

type StudentDBModel struct {
	ID           uuid.UUID      `db:"id"`
	UserID       uuid.UUID      `db:"user_id"`
	Age          int            `db:"age"`
	Classes      pq.StringArray `db:"classes_id"`
	Grades       pq.StringArray `db:"grades_id"`
	Disabilities pq.StringArray `db:"disabilities"`
	GuardianID   uuid.UUID      `db:"guardian_id"`
	Address      string         `db:"address"`
	Phone        string         `db:"phone"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    sql.NullTime   `db:"updated_at"`
	Active       bool           `db:"active"`
}

func (s *StudentDBModel) fromEntity(entity *studententities.Student) *StudentDBModel {
	disabilities := []string{}
	classes := []string{}
	grades := []string{}

	for _, d := range entity.Disabilities() {
		disabilities = append(disabilities, d.String())
	}

	for _, c := range entity.ClassesIDs() {
		classes = append(classes, c.String())
	}

	for _, g := range entity.GradesIDs() {
		grades = append(grades, g.String())
	}

	return &StudentDBModel{
		ID:           entity.ID(),
		UserID:       entity.UserID(),
		Age:          entity.Age(),
		Classes:      classes,
		Grades:       grades,
		Disabilities: disabilities,
		GuardianID:   entity.GuardianID(),
		Address:      entity.Address(),
		Phone:        entity.Phone(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    sql.NullTime{Time: entity.UpdatedAt()},
		Active:       entity.Active(),
	}
}

func (s *StudentDBModel) toEntity() *studententities.Student {
	disabilites := []enums.Disability{}
	classesIDs := []uuid.UUID{}
	gradesIDs := []uuid.UUID{}

	for _, d := range s.Disabilities {
		disabilites = append(disabilites, enums.ParseDisability(d))
	}

	for _, c := range s.Classes {
		classesIDs = append(classesIDs, uuid.MustParse(c))
	}

	for _, g := range s.Grades {
		gradesIDs = append(gradesIDs, uuid.MustParse(g))
	}

	return studententities.NewStudent(
		s.ID,
		s.UserID,
		s.Age,
		classesIDs,
		gradesIDs,
		disabilites,
		s.GuardianID,
		s.Address,
		s.Phone,
		s.CreatedAt,
		s.UpdatedAt.Time,
		s.Active,
	)

}

type CreateStudentDBModel struct {
	Id           uuid.UUID
	UserID       uuid.UUID `db:"user_id"`
	Age          int
	ClassesIDs   pq.StringArray `db:"classes_id"`
	GradesIDs    pq.StringArray `db:"grades_id"`
	Disabilities pq.StringArray
	GuardianID   uuid.UUID `db:"guardian_id"`
	Address      string
	Phone        string
}

func (c *CreateStudentDBModel) fromEntity(entity *studententities.CreateStudent) {
	disabilities := []string{}

	for _, d := range entity.Disabilities() {
		disabilities = append(disabilities, d.String())
	}

	c.Id = entity.ID()
	c.UserID = entity.UserID()
	c.Age = entity.Age()
	c.Disabilities = disabilities
	c.ClassesIDs = []string{}
	c.GradesIDs = []string{}
	c.GuardianID = entity.GuardianID()
	c.Address = entity.Address()
	c.Phone = entity.Phone()

}
