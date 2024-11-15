package contentrepository

import (
	"database/sql"

	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type ContentDbModel struct {
	ID            uuid.UUID    `db:"id"`
	Title         string       `db:"title"`
	Description   string       `db:"description"`
	ThumbnailLink string       `db:"thumbnail_link"`
	ContentLink   string       `db:"content_link"`
	ContentType   string       `db:"content_type"`
	CreatedAt     sql.NullTime `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
	Viewed        bool         `db:"viewed"`
}

func (c *ContentDbModel) fromEntity(entity *contententities.Content) {
	c.ID = entity.ID()
	c.Title = entity.Title()
	c.Description = entity.Description()
	c.ThumbnailLink = entity.ThumbnailLink()
	c.ContentLink = entity.ContentLink()
	c.ContentType = entity.ContentType().String()
	c.CreatedAt = sql.NullTime{Time: entity.CreatedAt(), Valid: true}
	c.UpdatedAt = sql.NullTime{Time: entity.UpdatedAt(), Valid: true}
	c.Viewed = entity.Viewed()
}

func (c *ContentDbModel) toEntity() *contententities.Content {
	return contententities.NewContent(
		c.ID,
		c.Title,
		c.Description,
		c.ThumbnailLink,
		c.ContentLink,
		enums.ParseContentType(c.ContentType),
		c.CreatedAt.Time,
		c.UpdatedAt.Time,
		c.Viewed,
	)
}

type CreateContentDbModel struct {
	ID            uuid.UUID `db:"id"`
	Title         string    `db:"title"`
	Description   string    `db:"description"`
	ThumbnailLink string    `db:"thumbnail_link"`
	ContentLink   string    `db:"content_link"`
	ContentType   string    `db:"content_type"`
}

func (c *CreateContentDbModel) fromEntity(entity *contententities.CreateContent) {
	c.ID = entity.ID()
	c.Title = entity.Title()
	c.Description = entity.Description()
	c.ThumbnailLink = entity.ThumbnailLink()
	c.ContentLink = entity.ContentLink()
	c.ContentType = entity.ContentType().String()
}
