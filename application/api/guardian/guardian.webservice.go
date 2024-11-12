package guardianweb

import (
	"github.com/google/uuid"
	guardianentities "github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
	userentities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	iservices "github.com/joaofilippe/edu-uni-srv/domain/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	guardianService iservices.IGuardianService
	userService     iservices.IUserService
}

func NewGuardianWeb(
	guardianService *iservices.IGuardianService,
	userService *iservices.IUserService,
) *WebService {
	return &WebService{
		*guardianService,
		*userService,
	}
}

func (sw *WebService) CreateGuardian(c echo.Context) error {
	req := new(CreateGuardianRequest)

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

	studentID, err := uuid.Parse(req.StudentID)
	if err != nil {
		return c.JSON(400, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error parsing student ID",
			err.Error(),
		})
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



	createGuardian := guardianentities.NewCreateGuardian(
		userID,
		studentID,
	)

	guardianID, err := sw.guardianService.Create(createGuardian)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error creating guardian",
			err.Error(),
		})
	}

	return c.JSON(200, struct {
		Message    string    `json:"message"`
		GuardianID uuid.UUID `json:"guardian_id"`
	}{
		"Guardian created",
		guardianID,
	})

}
