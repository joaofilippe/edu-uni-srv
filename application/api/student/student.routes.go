package studentweb

import (
	"github.com/joaofilippe/edu-uni-srv/common/tokenmanager"
	"github.com/labstack/echo/v4"
)

func (sw *WebService) BuildRoutes(server *echo.Echo) {
	student := server.Group("/student")

	student.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the Student API")
	})
	student.GET("/:id", sw.GetStudentByID)
	student.POST("", sw.CreateStudent, authAdminTeacherMiddleware)
	student.POST("/enroll", sw.EnrollStudent, authAdminTeacherMiddleware)
}

func authAdminTeacherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(401, struct {
				Message string `json:"message"`
			}{
				"Unauthorized",
			})
		}

		claims, err := tokenmanager.ValidateToken(token)
		if err != nil {
			return c.JSON(401, struct {
				Message string `json:"message"`
			}{
				"Unauthorized",
			})
		}

		if claims["userType"] != "Administrator" && claims["userType"] != "Teacher"{
			return c.JSON(401, struct {
				Message string `json:"message"`
			}{
				"Unauthorized",
			})
		}

		return next(c)
	}
}