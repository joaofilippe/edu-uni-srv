package teacher

type Teacher struct {
	id       int
	userId int
	name     string
	classIds []int
}

func NewTeacher(id int, name string, classIds []int) *Teacher {
	return &Teacher{id: id, name: name, classIds: classIds}
}

func (t *Teacher) ID() int {
	return t.id
}

func (t *Teacher) Name() string {
	return t.name
}

func (t *Teacher) ClassIds() []int {
	return t.classIds
}