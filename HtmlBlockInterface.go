package ui

type HtmlBlockInterface interface {
	BlockInterface

	// Serialization
	ToHTML() string
}
