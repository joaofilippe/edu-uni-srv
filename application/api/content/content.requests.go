package contentweb

type CreateContentRequest struct {
	Title         string `json:"title"`
	ClassID       string `json:"class_id"`
	Description   string `json:"description"`
	ThumbnailLink string `json:"thumbnail_link"`
	ContentLink   string `json:"content_link"`
	ContentType   string `json:"content_type"`
}
