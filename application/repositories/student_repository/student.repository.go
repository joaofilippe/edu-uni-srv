package studentrepository

import (
	"database/sql"

	"github.com/google/uuid"
	studentsEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type StudentRepository struct {
	conn *database.DBConnection
}

func NewStudentRepository(conn *database.DBConnection) *StudentRepository {
	return &StudentRepository{conn: conn}
}

func (s *StudentRepository) Save(user *studentsEntities.CreateStudent) error {
	return nil
}

func (s *StudentRepository) FindAll() ([]*studentsEntities.Student, error) {
	return nil, nil
}

func (s *StudentRepository) FindByID(id uuid.UUID) (*studentsEntities.Student, error) {
	return nil, nil
}

func (s *StudentRepository) FindByUserID(userID uuid.UUID) (*studentsEntities.Student, error) {
	studentDB := &StudentDBModel{}

	err := s.conn.DBConnection.Get(
		studentDB, 
		FindByUserIDQuery, 
		userID,
	)
	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, err
	}

	return studentDB.toEntity(), err
}

func (s *StudentRepository) FindByEmail(email string) (*studentsEntities.Student, error) {
	return nil, nil
}

func (s *StudentRepository) FindByUsername(username string) (*studentsEntities.Student, error) {
	return nil, nil
}

func (s *StudentRepository) Delete(id uuid.UUID) error {
	return nil
}