package classweb

import "github.com/labstack/echo/v4"

func (cw *WebService) BuildRoutes(server *echo.Echo) {
	class := server.Group("/class")

	class.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the Class API")
	})

	class.POST("", cw.CreateClass)
}