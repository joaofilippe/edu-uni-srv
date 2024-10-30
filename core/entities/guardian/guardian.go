package guardian

type Guardian struct {
	id       int
	userId   int
	name     string
	email    string
	password string
	studentId int
}

func NewGuardian(
	id int,
	userId int,
	name string,
	email string,
	password string,
	studentId int,
) *Guardian {
	return &Guardian{
		id,
		userId,
		name,
		email,
		password,
		studentId,
	}
}

func (g *Guardian) Id() int {
	return g.id
}

func (g *Guardian) UserId() int {
	return g.userId
}


func (g *Guardian) Name() string {
	return g.name
}

func (g *Guardian) Email() string {
	return g.email
}

func (g *Guardian) Password() string {
	return g.password
}

func (g *Guardian) StudentId() int {
	return g.studentId
}
