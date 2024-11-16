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
