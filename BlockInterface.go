package ui

type BlockInterface interface {
	Content() string
	Children() []*Block
	Parameters() map[string]string
	Parameter(key string) string
	SetParameter(key string, value string) *Block
	ToMap() map[string]interface{}
	ToJson() (string, error)
	ToJsonPretty() (string, error)
	ID() string
	SetID(id string) *Block
	Type() string
	SetType(blockType string) *Block
}
