package enums

type ContentType int

const (
	Audio ContentType = iota
	VideoSubtitled
	VideoWithSignLanguage
	NarratedText
	BrailleText
)

func ParseContentType(contentType string) ContentType {
	switch contentType {
	case "Audio":
		return Audio
	case "Video Subtitled", "video":
		return VideoSubtitled
	case "Video with Sign Language":
		return VideoWithSignLanguage
	case "Narrated Text":
		return NarratedText
	case "Braille Text":
		return BrailleText
	default:
		return Audio
	}
}

func (ct ContentType) String() string {
	return [...]string{"Audio", "Video Subtitled", "Video with Sign Language", "Narrated Text", "Braille Text"}[ct]
}