package ui

import (
	"encoding/json"
	"reflect"

	"github.com/gouniverse/uid"
)

// == CONSTUCTORS ==============================================================

// NewBlock returns a new block instance, and sets the default ID
func NewBlock() BlockInterface {
	block := &Block{}

	block.SetID(uid.HumanUid()) // default, can be changed by SetID if needed

	return block
}

func NewBlockFromJson(blockJson string) (BlockInterface, error) {
	blockMap := map[string]any{}

	err := json.Unmarshal([]byte(blockJson), &blockMap)

	if err != nil {
		return nil, err
	}

	return ConvertMapToBlock(blockMap)
}

// BlockFromMap creates a block from a map
func NewBlockFromMap(m map[string]any) BlockInterface {
	id := ""

	if idMap, ok := m["id"].(string); ok {
		id = idMap
	}

	blockType := ""

	if blockTypeMap, ok := m["type"].(string); ok {
		blockType = blockTypeMap
	}

	parameters := map[string]string{}

	if parametersMap, ok := m["parameters"].(map[string]string); ok {
		for k, v := range parametersMap {
			parameters[k] = v
		}
	}

	children := []BlockInterface{}

	if childrenAny, ok := m["children"]; ok {
		typeOfChildren := reflect.TypeOf(childrenAny).Elem()

		kind := typeOfChildren.Kind()

		if kind == reflect.Interface {
			childrenMap := childrenAny.([]BlockInterface)
			children = childrenMap
		}

		if kind == reflect.Map {
			childrenMap := childrenAny.([]map[string]any)
			for _, c := range childrenMap {
				child := NewBlockFromMap(c)
				if child == nil {
					continue
				}
				children = append(children, child)
			}
		}
	}

	block := NewBlock()
	block.SetID(id)
	block.SetType(blockType)
	block.SetParameters(parameters)
	block.SetChildren(children)
	return block
}

// == TYPE ====================================================================

type Block struct {
	id         string
	blockType  string
	children   []BlockInterface
	parameters map[string]string
}

// type BlockConfig struct {
// 	ID         string
// 	Type       string
// 	Parameters map[string]string
// 	Children   []BlockInterface
// }

// == INTERFACE IMPLEMENTATION ================================================

func (b *Block) AddChild(child BlockInterface) {
	if b.children == nil {
		b.children = []BlockInterface{}
	}
	b.children = append(b.children, child)
}

func (b *Block) AddChildren(children []BlockInterface) {
	if b.children == nil {
		b.children = []BlockInterface{}
	}
	b.children = append(b.children, children...)
}

func (b *Block) Children() []BlockInterface {
	return b.children
}

func (b *Block) SetChildren(children []BlockInterface) {
	b.children = children
}

func (b *Block) ID() string {
	return b.id
}

func (b *Block) SetID(id string) {
	b.id = id
}

func (b *Block) HasParameter(key string) bool {
	_, ok := b.parameters[key]
	return ok
}

func (b *Block) Parameters() map[string]string {
	return b.parameters
}

func (b *Block) SetParameters(parameters map[string]string) {
	b.parameters = parameters
}

func (b *Block) Parameter(key string) string {
	if value, ok := b.parameters[key]; ok {
		return value
	}

	return ""
}

func (b *Block) SetParameter(key, value string) {
	if b.parameters == nil {
		b.parameters = map[string]string{}
	}
	b.parameters[key] = value
}

func (b *Block) Type() string {
	return b.blockType
}

func (b *Block) SetType(blockType string) {
	b.blockType = blockType
}

func (b *Block) ToMap() map[string]interface{} {
	childrenMap := []map[string]interface{}{}

	for _, child := range b.children {
		childrenMap = append(childrenMap, child.ToMap())
	}

	return map[string]any{
		"id":         b.ID(),
		"type":       b.Type(),
		"parameters": b.Parameters(),
		"children":   childrenMap,
	}
}

func (b *Block) ToJson() (string, error) {
	jsonObject := b.ToJsonObject()

	jsonBytes, err := json.Marshal(jsonObject)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (b *Block) ToJsonPretty() (string, error) {
	jsonObject := b.ToJsonObject()

	jsonBytes, err := json.MarshalIndent(jsonObject, "", "  ")

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (b *Block) ToJsonObject() blockJsonObject {
	parameters := b.Parameters()
	if len(parameters) < 1 {
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
		Parameters: parameters,
		Children:   childrenJsonObject,
	}
}

type blockJsonObject struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Content    string            `json:"content"`
	Parameters map[string]string `json:"parameters"`
	Children   []blockJsonObject `json:"children"`
}
