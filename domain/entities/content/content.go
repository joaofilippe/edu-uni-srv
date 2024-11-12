package contententities

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type Content struct {
	id            uuid.UUID
	title         string
	description   string
	thumbnailLink string
	contentLink   string
	contentType   enums.ContentType
}

func NewContent(
	id uuid.UUID,
	title string,
	description string,
	thumbnailLink string,
	contentLink string,
	contentType enums.ContentType,
) *Content {
	content := &Content{
		id,
		title,
		description,
		thumbnailLink,
		contentLink,
		contentType,
	}

	if !content.validateContentLink() {
		return nil
	}

	return content
}

func (c *Content) ID() uuid.UUID {
	return c.id
}

func (c *Content) Title() string {
	return c.title
}

func (c *Content) Description() string {
	return c.description
}

func (c *Content) Link() string {
	return c.contentLink
}

func (c *Content) ContentType() enums.ContentType {
	return c.contentType
}

func (c *Content) ThumbnailLink() string {
	return c.thumbnailLink
}

func (c *Content) validateContentLink() bool {
	return c.contentLink != ""
}
