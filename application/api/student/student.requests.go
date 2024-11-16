package studentweb

type CreateStudentRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Disabilities []string `json:"disabilities"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type EnrollStudentRequest struct {
	ClassID string `json:"class_id"`
	StudentID string `json:"student_id"`
}
