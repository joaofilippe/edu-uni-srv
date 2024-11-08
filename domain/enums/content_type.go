package enums

type ContentType int

const (
	Audio ContentType = iota
	VideoSubtitled
	VideoWithSignLanguage
	NarratedText
	BrailleText
)

func (ct ContentType) String() string {
	return [...]string{"Audio", "Video Subtitled", "Video with Sign Language", "Narrated Text", "Braille Text"}[ct]
}