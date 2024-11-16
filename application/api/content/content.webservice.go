package contentweb

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	iservices "github.com/joaofilippe/edu-uni-srv/domain/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	contentService iservices.IContentService
}

func NewContentWeb(contentService *iservices.IContentService) *WebService {
	return &WebService{
		contentService: *contentService,
	}
}

func (cw *WebService) CreateContent(c echo.Context) error {
	req := new(CreateContentRequest)
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

	contentType := enums.ParseContentType(req.ContentType)
	classID, err := uuid.Parse(req.ClassID)
	if err != nil {
		return c.JSON(400, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error parsing class ID",
			err.Error(),
		})
	}

	newContent := contententities.NewCreateContent(
		req.Title,
		classID,
		req.Description,
		req.ThumbnailLink,
		req.ContentLink,
		contentType,
	)

	contentID, err := cw.contentService.Create(newContent)
	if err != nil {
		return c.JSON(500, struct {
			Message string `json:"message"`
			Err     string `json:"error"`
		}{
			"Error creating content",
			err.Error(),
		})
	}

	return c.JSON(200, struct {
		Message   string `json:"message"`
		ContentID string `json:"content_id"`
	}{
		"Content created successfully",
		contentID.String(),
	})
}
