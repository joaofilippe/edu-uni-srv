package guardian

type Guardian struct {
	id       int
	name     string
	email    string
	password string
	studentId int
}

func NewGuardian(
	id int,
	name string,
	email string,
	password string,
	studentId int,
) *Guardian {
	return &Guardian{
		id,
		name,
		email,
		password,
		studentId,
	}
}

func (g *Guardian) ID() int {
	return g.id
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

func (g *Guardian) StudentID() int {
	return g.studentId
}