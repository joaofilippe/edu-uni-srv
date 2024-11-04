package contententities

import "github.com/joaofilippe/edu-uni-srv/core/enums"

type Content struct {
	id            int
	title         string
	description   string
	thumbnailLink string
	contentLink   string
	contentType   enums.ContentType
	// TODO campo para indicar o tipo de deficiência do aluno
}

func NewContent(
	id int,
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
