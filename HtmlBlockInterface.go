package ui

type HtmlBlockInterface interface {
	BlockInterface

	// Serialization
	ToHtml() string
}
