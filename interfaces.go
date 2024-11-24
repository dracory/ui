package ui

type BlockInterface interface {
	IDInterface
	ChildrenInterface
	ParametersInterface
	TypeInterface

	// Serialization
	ToMapInterface
	ToJsonInterface
	ToJsonObjectInterface
	ToJsonPrettyInterface
}

type ChildrenInterface interface {
	Children() []BlockInterface
	SetChildren([]BlockInterface)
	AddChild(BlockInterface)
	AddChildren([]BlockInterface)
}

type IDInterface interface {
	ID() string
	SetID(string)
}

// type HtmlBlockInterface interface {
// 	BlockInterface
// 	ToHTMLInterface
// }

type ParametersInterface interface {
	HasParameter(key string) bool
	Parameter(key string) string
	SetParameter(key string, value string)
	Parameters() map[string]string
	SetParameters(map[string]string)
}

type TypeInterface interface {
	Type() string
	SetType(string)
}

type ToHTMLInterface interface {
	ToHTML() string
}

type ToJsonInterface interface {
	ToJson() (string, error)
}

type ToJsonPrettyInterface interface {
	ToJsonPretty() (string, error)
}

type ToJsonObjectInterface interface {
	ToJsonObject() blockJsonObject
}

type ToMapInterface interface {
	ToMap() map[string]interface{}
}

type parametersImplementation struct {
	parameters map[string]string
}

var _ ParametersInterface = (*parametersImplementation)(nil)

func (b *parametersImplementation) HasParameter(key string) bool {
	_, ok := b.parameters[key]
	return ok
}

func (b *parametersImplementation) Parameters() map[string]string {
	return b.parameters
}

func (b *parametersImplementation) SetParameters(parameters map[string]string) {
	b.parameters = parameters
}

func (b *parametersImplementation) Parameter(key string) string {
	if value, ok := b.parameters[key]; ok {
		return value
	}

	return ""
}

func (b *parametersImplementation) SetParameter(key, value string) {
	if b.parameters == nil {
		b.parameters = map[string]string{}
	}
	b.parameters[key] = value
}

type BlockBuilderInterface interface {
	WithID(string) BlockBuilderInterface
	WithType(string) BlockBuilderInterface
	WithParameters(map[string]string) BlockBuilderInterface
	WithChildren([]BlockInterface) BlockBuilderInterface
	Build() BlockInterface
}

type blockBuilder struct {
	id         string
	blockType  string
	parameters map[string]string
	children   []BlockInterface
}

var _ BlockBuilderInterface = (*blockBuilder)(nil)

func NewBlockBuilder() BlockBuilderInterface {
	return &blockBuilder{parameters: map[string]string{}}
}

func (b *blockBuilder) WithID(id string) BlockBuilderInterface {
	b.id = id
	return b
}

func (b *blockBuilder) WithType(blockType string) BlockBuilderInterface {
	b.blockType = blockType
	return b
}

func (b *blockBuilder) WithParameters(parameters map[string]string) BlockBuilderInterface {
	b.parameters = parameters
	return b
}

func (b *blockBuilder) WithChildren(children []BlockInterface) BlockBuilderInterface {
	b.children = children
	return b
}

func (b *blockBuilder) Build() BlockInterface {
	if b.parameters == nil {
		b.parameters = map[string]string{}
	}

	if b.children == nil {
		b.children = []BlockInterface{}
	}

	block := NewBlock()
	block.SetID(b.id)
	block.SetType(b.blockType)
	block.SetParameters(b.parameters)
	block.SetChildren(b.children)
	return block
}
