package contentweb

type CreateContentRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	ThumbnailLink string `json:"thumbnail_link"`
	ContentLink   string `json:"content_link"`
	ContentType   string `json:"content_type"`
}