package teacherrepository

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type TeacherRepository struct {
	conn *database.DBConnection
}

func NewTeacherRepository(conn *database.DBConnection) *TeacherRepository {
	return &TeacherRepository{conn: conn}
}

func (t *TeacherRepository) Save(teacher *teacherentities.CreateTeacher) error {
	createTeacherDB := &CreateTeacherDbModel{}
	createTeacherDB.fromEntity(teacher)

	tx := t.conn.DBConnection.MustBegin()
	_, err := tx.NamedExec(Save, createTeacherDB)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (t *TeacherRepository) FindAll() ([]*teacherentities.Teacher, error) {
	return nil, nil
}

func (t *TeacherRepository) FindByID(id uuid.UUID) (*teacherentities.Teacher, error) {
	return nil, nil
}

func (t *TeacherRepository) FindByUserID(userID uuid.UUID) (*teacherentities.Teacher, error) {
	teacherDB := &TeacherDBModel{}
	err := t.conn.DBConnection.Get(teacherDB, FindByUserID, userID)
	if err != nil {
		return nil, err
	}

	return teacherDB.toEntity(), nil
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
