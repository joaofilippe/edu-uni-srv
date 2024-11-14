package contententities

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type CreateContent struct {
	id            uuid.UUID
	title         string
	description   string
	thumbnailLink string
	contentLink   string
	contentType   enums.ContentType
}

func NewCreateContent(
	id uuid.UUID,
	title string,
	description string,
	thumbnailLink string,
	contentLink string,
	contentType enums.ContentType,
) *CreateContent {
	return &CreateContent{
		id,
		title,
		description,
		thumbnailLink,
		contentLink,
		contentType,
	}
}

func (c *CreateContent) ID() uuid.UUID {
	return c.id
}

func (c *CreateContent) Title() string {
	return c.title
}

func (c *CreateContent) Description() string {
	return c.description
}

func (c *CreateContent) ThumbnailLink() string {
	return c.thumbnailLink
}

func (c *CreateContent) ContentLink() string {
	return c.contentLink
}

func (c *CreateContent) ContentType() enums.ContentType {
	return c.contentType
}

func (c *CreateContent) validateContentLink() bool {
	return true
}

func (c *CreateContent) SetID(id uuid.UUID) {
	c.id = id
}
