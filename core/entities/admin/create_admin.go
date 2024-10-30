package admin

type CreateAdmin struct {
	name     string
	email    string
	password string
}

func NewCreateAdmin(
	name string,
	email string,
	password string,
) *CreateAdmin {
	return &CreateAdmin{
		name,
		email,
		password,
	}
}

func (c *CreateAdmin) Name() string {
	return c.name
}

func (c *CreateAdmin) Email() string {
	return c.email
}

func (c *CreateAdmin) Password() string {
	return c.password
}
