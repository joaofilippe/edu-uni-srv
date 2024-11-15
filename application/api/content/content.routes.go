package contentweb

import "github.com/labstack/echo/v4"

func (cw *WebService) BuildRoutes(server *echo.Echo) {
	content := server.Group("/content")

	content.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the Content API")
	})

	content.POST("", cw.CreateContent)
}