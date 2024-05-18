package ui

import "encoding/json"

func BlocksFromJson(blocksJson string) ([]Block, error) {
	blocksMap := []map[string]any{}

	err := json.Unmarshal([]byte(blocksJson), &blocksMap)

	if err != nil {
		return nil, err
	}

	blocks := []Block{}

	for _, blockMap := range blocksMap {
		block, err := mapStringAnyToBlock(blockMap)

		if err != nil {
			return nil, err
		}

		blocks = append(blocks, *block)
	}

	return blocks, nil
}
