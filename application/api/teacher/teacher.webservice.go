package teacherweb

import (
	"github.com/google/uuid"
	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	userentities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	iservices "github.com/joaofilippe/edu-uni-srv/domain/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	teacherService iservices.ITeacherService
	userService    iservices.IUserService
}

func NewTeacherWeb(
	studentService *iservices.ITeacherService,
	userService *iservices.IUserService,
) *WebService {
	return &WebService{
		userService:    *userService,
		teacherService: *studentService,
	}
}

func (sw *WebService) CreateTeacher(c echo.Context) error {
	req := new(CreateTeacherRequest)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(400, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error parsing request",
			err.Error(),
		})
	}

	createUser := userentities.NewCreateUser(
		req.Name,
		req.Password,
		req.Email,
		enums.Teacher,
		nil,
	)

	userID, err := sw.userService.Create(createUser)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error creating user",
			err.Error(),
		})
	}

	newTeacher := teacherentities.NewCreateTeacher(
		userID,
	)

	teacherID, err := sw.teacherService.Create(newTeacher)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error creating teacher",
			err.Error(),
		})

	}

	return c.JSON(200, struct {
		Message string    `json:"message"`
		ID      uuid.UUID `json:"id"`
	}{
		"Teacher created",
		teacherID,
	})

}
