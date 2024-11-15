package contententities

import (
	"time"

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
	createdAt     time.Time
	updatedAt     time.Time
	viewed bool
}

func NewContent(
	id uuid.UUID,
	title string,
	description string,
	thumbnailLink string,
	contentLink string,
	contentType enums.ContentType,
	createdAt time.Time,
	updatedAt time.Time,
	viewed bool,
) *Content {
	content := &Content{
		id,
		title,
		description,
		thumbnailLink,
		contentLink,
		contentType,
		createdAt,
		updatedAt,
		viewed,
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

func (c *Content) ContentLink() string {
	return c.contentLink
}

func (c *Content) ContentType() enums.ContentType {
	return c.contentType
}

func (c *Content) ThumbnailLink() string {
	return c.thumbnailLink
}

func (c *Content) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Content) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Content) Viewed() bool {
	return c.viewed
}

func (c *Content) validateContentLink() bool {
	return c.contentLink != ""
}
