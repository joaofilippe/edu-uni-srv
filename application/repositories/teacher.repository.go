package repositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type TeacherRepository struct {
	conn *database.DBConnection
}

func NewTeacherRepository(conn *database.DBConnection) *TeacherRepository {
	return &TeacherRepository{conn: conn}
}

func (t *TeacherRepository) Save(teacher *teacherentities.CreateTeacher) error {
	return nil
}

func (t *TeacherRepository) FindAll() ([]*teacherentities.Teacher, error) {
	return nil, nil
}

func (t *TeacherRepository) FindByID(id uuid.UUID) (*teacherentities.Teacher, error) {
	return nil, nil
}

func (t *TeacherRepository) FindByEmail(email string) (*teacherentities.Teacher, error) {
	return nil, nil
}

func (t *TeacherRepository) FindByUsername(username string) (*teacherentities.Teacher, error) {
	return nil, nil
}

func (t *TeacherRepository) Update(teacher *teacherentities.Teacher) error {
	return nil
}

func (t *TeacherRepository) Delete(id uuid.UUID) error {
	return nil
}
