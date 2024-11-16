package ui

import (
	"encoding/json"
)

func BlockFromJson(blockJson string) (BlockInterface, error) {
	blockMap := map[string]any{}

	err := json.Unmarshal([]byte(blockJson), &blockMap)

	if err != nil {
		return nil, err
	}

	return mapToBlock(blockMap)
}
