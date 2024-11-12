package studentweb

import (
	"github.com/google/uuid"
	studententities "github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	userentities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	iservices "github.com/joaofilippe/edu-uni-srv/domain/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	studentService iservices.IStudentService
	userService    iservices.IUserService
}

func NewStudentWeb(
	studentService *iservices.IStudentService,
	userService *iservices.IUserService,
) *WebService {
	return &WebService{
		userService:    *userService,
		studentService: *studentService,
	}
}

func (sw *WebService) CreateStudent(c echo.Context) error {
	req := new(CreateStudentRequest)
	err := c.Bind(req)

	if err != nil {
		return c.JSON(400, "Invalid request")
	}

	createUser := userentities.NewCreateUser(
		req.Name,
		req.Password,
		req.Email,
		enums.Student,
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

	disabilites := []enums.Disability{}
	for _, d := range req.Disabilities {
		disabilites = append(disabilites, enums.ParseDisability(d))
	}

	newStudent := studententities.NewCreateStudent(
		userID,
		req.Age,
		disabilites,
		uuid.UUID{},
		req.Address,
		req.Phone,
	)

	_, err = sw.studentService.Create(newStudent)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error creating student",
			err.Error(),
		})
	}
	return c.JSON(201, struct {
		Message string `json:"message"`
	}{
		"Student created",
	})
}

func (sw *WebService) GetStudentByID(c echo.Context) error {
	param := c.Param("id")
	studentID := uuid.MustParse(param)
	result, err := sw.studentService.FindByUserID(studentID)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error getting student",
			err.Error(),
		})
	}

	student := new(StudentReponse)
	student.fromEntity(result)
	return c.JSON(200, struct {
		Message string          `json:"message"`
		Student *StudentReponse `json:"student"`
	}{
		"Student found",
		student,
	})
}
