package contententities

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type CreateContent struct {
	id            uuid.UUID
	classID       uuid.UUID
	title         string
	description   string
	thumbnailLink string
	contentLink   string
	contentType   enums.ContentType
}

func NewCreateContent(
	title string,
	classID uuid.UUID,
	description string,
	thumbnailLink string,
	contentLink string,
	contentType enums.ContentType,
) *CreateContent {
	return &CreateContent{
		title:         title,
		description:   description,
		thumbnailLink: thumbnailLink,
		contentLink:   contentLink,
		contentType:   contentType,
	}
}

func (c *CreateContent) ID() uuid.UUID {
	return c.id
}

func (c *CreateContent) ClassID() uuid.UUID {
	return c.classID
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
