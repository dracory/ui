package ui

type BlockInterface interface {
	// Content
	Content() string
	SetContent() BlockInterface

	// Children
	Children() []BlockInterface
	SetChildren(children []BlockInterface) BlockInterface

	// ID
	ID() string
	SetID(id string) *Block

	// Parameters
	Parameters() map[string]string
	Parameter(key string) string
	SetParameter(key string, value string) BlockInterface

	// Type
	Type() string
	SetType(blockType string) BlockInterface

	// Serialization
	ToMap() map[string]interface{}
	ToJson() (string, error)
	ToJsonPretty() (string, error)
}
