package admin

type Admin struct {
	id int
	userId int
	name string
	email string
	password string
}

func NewAdmin(
	id int,
	userId int,
	name string,
	email string,
	password string,
) *Admin {
	return &Admin{
		id,
		userId,
		name,
		email,
		password,
	}
}

func (a *Admin) ID() int {
	return a.id
}

func (a *Admin) Name() string {
	return a.name
}

func (a *Admin) Email() string {
	return a.email
}

func (a *Admin) Password() string {
	return a.password
}

func (a *Admin) UserId() int {
	return a.userId
}
