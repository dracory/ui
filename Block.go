package ui

import (
	"encoding/json"

	"github.com/gouniverse/uid"
)

var _ BlockInterface = (*Block)(nil) // verify it extends the interface

func NewBlock() *Block {
	block := &Block{}
	block.SetID(uid.HumanUid())
	return block
}

type Block struct {
	id         string
	blockType  string
	content    string
	children   []BlockInterface
	parameters map[string]string
}

func (b *Block) AddChild(child BlockInterface) BlockInterface {
	if b.children == nil {
		b.children = []BlockInterface{}
	}
	b.children = append(b.children, child)
	return b
}

func (b *Block) AddChildren(children []BlockInterface) BlockInterface {
	if b.children == nil {
		b.children = []BlockInterface{}
	}
	b.children = append(b.children, children...)
	return b
}

func (b *Block) Children() []BlockInterface {
	return b.children
}

func (b *Block) SetChildren(children []BlockInterface) BlockInterface {
	b.children = children
	return b
}

func (b *Block) Content() string {
	return b.content
}

func (b *Block) SetContent(content string) BlockInterface {
	b.content = content
	return b
}

func (b *Block) ID() string {
	return b.id
}

func (b *Block) SetID(id string) *Block {
	b.id = id
	return b
}

func (b *Block) Parameters() map[string]string {
	return b.parameters
}

func (b *Block) SetParameters(parameters map[string]string) BlockInterface {
	b.parameters = parameters
	return b
}

func (b *Block) Parameter(key string) string {
	if value, ok := b.parameters[key]; ok {
		return value
	}

	return ""
}

func (b *Block) SetParameter(key string, value string) BlockInterface {
	if b.parameters == nil {
		b.parameters = map[string]string{}
	}
	b.parameters[key] = value
	return b
}

func (b *Block) Type() string {
	return b.blockType
}

func (b *Block) SetType(blockType string) BlockInterface {
	b.blockType = blockType
	return b
}

func (b *Block) ToMap() map[string]interface{} {
	childrenMap := []map[string]interface{}{}

	for _, child := range b.children {
		childrenMap = append(childrenMap, child.ToMap())
	}

	return map[string]interface{}{
		"id":         b.ID(),
		"type":       b.Type(),
		"content":    b.Content(),
		"parameters": b.Parameters(),
		"children":   childrenMap,
	}
}

func (b *Block) ToJson() (string, error) {
	jsonObject := b.ToJsonObject()

	json, err := json.Marshal(jsonObject)

	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (b *Block) ToJsonPretty() (string, error) {
	jsonObject := b.ToJsonObject()

	json, err := json.MarshalIndent(jsonObject, "", "  ")

	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (b *Block) ToJsonObject() blockJsonObject {
	parameters := b.Parameters()
	if parameters == nil || len(parameters) < 1 {
		parameters = make(map[string]string)
	}

	childrenJsonObject := make([]blockJsonObject, 0)

	for _, child := range b.Children() {
		childBlock := child.(*Block)
		childrenJsonObject = append(childrenJsonObject, childBlock.ToJsonObject())
	}

	return blockJsonObject{
		ID:         b.ID(),
		Type:       b.Type(),
		Content:    b.Content(),
		Parameters: parameters,
		Children:   childrenJsonObject,
	}
}
