package classweb

type CreateClassRequest struct {
	Name        string   `json:"name"`
	TeacherID   string   `json:"teacher_id"`
}