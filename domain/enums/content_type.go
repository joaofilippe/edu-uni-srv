package enums

import "strings"

type ContentType int

const (
	Audio ContentType = iota
	SubtitledText
	LibrasVideo
	AudioText
	BrailleText
)

func ParseContentType(contentType string) ContentType {
	switch strings.ToLower(contentType) {
	case "audio":
		return Audio
	case "subtitled-video", "video":
		return SubtitledText
	case "libras-video":
		return LibrasVideo
	case "audio-text":
		return AudioText
	case "braille-text":
		return BrailleText
	default:
		return Audio
	}
}

func (ct ContentType) String() string {
	return [...]string{"audio", "subtitled-video", "libras-video", "audio-text", "braille-text"}[ct]
}