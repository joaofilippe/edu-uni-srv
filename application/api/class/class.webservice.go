package classweb

import (
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
	return nil
}
