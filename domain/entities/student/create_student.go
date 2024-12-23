package studententities

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type CreateStudent struct {
	id           uuid.UUID
	userID       uuid.UUID
	age          int
	disabilities []enums.Disability
	guardianID   uuid.UUID
	address      string
	phone        string
}

func NewCreateStudent(
	userID uuid.UUID,
	age int,
	disabilities []enums.Disability,
	guardianID uuid.UUID,
	address string,
	phone string,
) *CreateStudent {
	return &CreateStudent{
		uuid.UUID{},
		userID,
		age,
		disabilities,
		guardianID,
		address,
		phone,
	}
}

func (cs *CreateStudent) ID() uuid.UUID {
	return cs.id
}

func (cs *CreateStudent) UserID() uuid.UUID {
	return cs.userID
}


func (cs *CreateStudent) Age() int {
	return cs.age
}

func (cs *CreateStudent) Disabilities() []enums.Disability {
	return cs.disabilities
}

func (cs *CreateStudent) GuardianID() uuid.UUID {
	return cs.guardianID
}

func (cs *CreateStudent) Address() string {
	return cs.address
}

func (cs *CreateStudent) Phone() string {
	return cs.phone
}

func (cs *CreateStudent) SetId(id uuid.UUID) {
	cs.id = id
}