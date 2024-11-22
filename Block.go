package ui

import (
	"encoding/json"
	"reflect"

	"github.com/gouniverse/uid"
)

// == CONSTUCTORS ==============================================================

// Block returns a new block instance, and sets default ID
func Block() BlockInterface {
	block := &block{}

	block.SetID(uid.HumanUid()) // default, can be changed by SetID if needed

	return block
}

func BlockFromJson(blockJson string) (BlockInterface, error) {
	blockMap := map[string]any{}

	err := json.Unmarshal([]byte(blockJson), &blockMap)

	if err != nil {
		return nil, err
	}

	return mapToBlock(blockMap)
}

// BlockFromMap creates a block from a map
func BlockFromMap(m map[string]interface{}) BlockInterface {
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
			childrenMap := childrenAny.([]map[string]interface{})
			for _, c := range childrenMap {
				child := BlockFromMap(c)
				if child == nil {
					continue
				}
				children = append(children, child)
			}
		}
	}

	block := Block()
	block.SetID(id)
	block.SetType(blockType)
	block.SetParameters(parameters)
	block.SetChildren(children)
	return block
}

// == TYPE ====================================================================

type block struct {
	id         string
	blockType  string
	children   []BlockInterface
	parameters map[string]string
}

type BlockConfig struct {
	ID         string
	Type       string
	Parameters map[string]string
	Children   []BlockInterface
}

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

type blockJsonObject struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Content    string            `json:"content"`
	Parameters map[string]string `json:"parameters"`
	Children   []blockJsonObject `json:"children"`
}
