package ui

import (
	"encoding/json"

	"github.com/gouniverse/uid"
)

// == TYPE ====================================================================

type block struct {
	id        string
	blockType string
	// content    string
	children   []BlockInterface
	parameters map[string]string
}

// == CONSTUCTOR ==============================================================

type BlockConfig struct {
	ID         string
	Type       string
	Parameters map[string]string
	Children   []BlockInterface
}

// Block returns a new block instance, and sets default ID
func Block() BlockInterface {
	block := &block{}

	block.SetID(uid.HumanUid()) // default, can be changed by SetID if needed

	return block
}

// == INTERFACE VERIFICATION ==================================================

var _ BlockInterface = (*block)(nil) // verify it extends the interface

// == INTERFACE IMPLEMENTATION ================================================

func (b *block) AddChild(child BlockInterface) {
	if b.children == nil {
		b.children = []BlockInterface{}
	}
	b.children = append(b.children, child)
}

func (b *block) AddChildren(children []BlockInterface) {
	if b.children == nil {
		b.children = []BlockInterface{}
	}
	b.children = append(b.children, children...)
}

func (b *block) Children() []BlockInterface {
	return b.children
}

func (b *block) SetChildren(children []BlockInterface) {
	b.children = children
}

func (b *block) ID() string {
	return b.id
}

func (b *block) SetID(id string) {
	b.id = id
}

func (b *block) Parameters() map[string]string {
	return b.parameters
}

func (b *block) SetParameters(parameters map[string]string) {
	b.parameters = parameters
}

func (b *block) Parameter(key string) string {
	if value, ok := b.parameters[key]; ok {
		return value
	}

	return ""
}

func (b *block) SetParameter(key string, value string) {
	if b.parameters == nil {
		b.parameters = map[string]string{}
	}
	b.parameters[key] = value
}

func (b *block) Type() string {
	return b.blockType
}

func (b *block) SetType(blockType string) {
	b.blockType = blockType
}

func (b *block) ToMap() map[string]interface{} {
	childrenMap := []map[string]interface{}{}

	for _, child := range b.children {
		childrenMap = append(childrenMap, child.ToMap())
	}

	return map[string]interface{}{
		"id":         b.ID(),
		"type":       b.Type(),
		"parameters": b.Parameters(),
		"children":   childrenMap,
	}
}

func (b *block) ToJson() (string, error) {
	jsonObject := b.ToJsonObject()

	json, err := json.Marshal(jsonObject)

	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (b *block) ToJsonPretty() (string, error) {
	jsonObject := b.ToJsonObject()

	json, err := json.MarshalIndent(jsonObject, "", "  ")

	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (b *block) ToJsonObject() blockJsonObject {
	parameters := b.Parameters()
	if parameters == nil || len(parameters) < 1 {
		parameters = make(map[string]string)
	}

	childrenJsonObject := make([]blockJsonObject, 0)

	for _, child := range b.Children() {
		childBlock := child.(*block)
		childrenJsonObject = append(childrenJsonObject, childBlock.ToJsonObject())
	}

	return blockJsonObject{
		ID:   b.ID(),
		Type: b.Type(),
		// Content:    b.Content(),
		Parameters: parameters,
		Children:   childrenJsonObject,
	}
}
