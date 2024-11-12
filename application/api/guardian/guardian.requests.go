package guardianweb

type CreateGuardianRequest struct {
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
