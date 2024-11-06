package adminweb

import "github.com/labstack/echo/v4"

func (aw *WebService) BuildRoutes(server *echo.Echo) {
	admin := server.Group("/admin")

	admin.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the Admin API")
	})
	admin.POST("", aw.CreateAdmin)
}
