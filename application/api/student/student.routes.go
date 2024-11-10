package studentweb

import "github.com/labstack/echo/v4"

func (sw *WebService) BuildRoutes(server *echo.Echo) {
	student := server.Group("/student")

	student.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the Student API")
	})
	student.POST("", sw.CreateStudent)
}