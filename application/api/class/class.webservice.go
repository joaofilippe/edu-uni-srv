package classweb

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/class"
	iservices "github.com/joaofilippe/edu-uni-srv/domain/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	classService iservices.IClassService
}

func NewClassWeb(classService *iservices.IClassService) *WebService {
	return &WebService{
		classService: *classService,
	}
}

func (cw *WebService) CreateClass(c echo.Context) error {
	req := new(CreateClassRequest)
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

	teacherID, err := uuid.Parse(req.TeacherID)
	if err != nil {
		return c.JSON(400, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error parsing teacher ID",
			err.Error(),
		})
	}

	createClass := classentities.NewCreateClass(
		req.Name,
		teacherID,
	)

	classID, err := cw.classService.Create(createClass)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error creating class",
			err.Error(),
		})
	}

	return c.JSON(201, struct {
		Message string `json:"message"`
		ClassID string `json:"class_id"`
	}{
		"Class created",
		classID.String(),
	})

}
