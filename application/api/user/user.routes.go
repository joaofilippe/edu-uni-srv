package userweb

import "github.com/labstack/echo/v4"

func (uw *WebService) BuildRoutes(server *echo.Echo) {
	user := server.Group("/user")

	user.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the User API")
	})
	user.POST("", uw.CreateUser)
	user.POST("/login", uw.Login)
}
