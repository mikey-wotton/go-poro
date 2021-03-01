package go_poro

type Tag struct {
	Name        string
	Description string
	Colour      TagColour
}

type TagColour string

const (
	TagColourUnknown TagColour = ""
	TagColourGreen   TagColour = "green"
	TagColourYellow  TagColour = "yellow"
	TagColourRed     TagColour = "red"
)
