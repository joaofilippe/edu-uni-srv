package content

import "github.com/joaofilippe/edu-uni-srv/core/enums"

type Content struct {
	id            int
	title         string
	description   string
	link          string
	contentType enums.ContentType
}

func NewContent(
	id int,
	title string,
	description string,
	link string,
	contentType enums.ContentType,
) *Content {
	return &Content{
		id,
		title,
		description,
		link,
		contentType,
	}
}

func (c *Content) ID() int {
	return c.id
}

func (c *Content) Title() string {
	return c.title
}

func (c *Content) Description() string {
	return c.description
}

func (c *Content) Link() string {
	return c.link
}

func (c *Content) ContentType() enums.ContentType {
	return c.contentType
}
