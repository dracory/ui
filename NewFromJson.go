package ui

import (
	"encoding/json"
)

func NewFromJson(blockJson string) (BlockInterface, error) {
	blockMap := map[string]any{}

	err := json.Unmarshal([]byte(blockJson), &blockMap)

	if err != nil {
		return nil, err
	}

	return mapStringAnyToBlock(blockMap)
}

func mapStringAnyToBlock(blockMap map[string]any) (BlockInterface, error) {
	blockMap, err := mapStringAnyToBlockMap(blockMap)

	if err != nil {
		return nil, err
	}

	return NewFromMap(blockMap), nil
}
