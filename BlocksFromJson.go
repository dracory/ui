package ui

import "encoding/json"

func BlocksFromJson(blocksJson string) ([]BlockInterface, error) {
	blocksMap := []map[string]any{}

	err := json.Unmarshal([]byte(blocksJson), &blocksMap)

	if err != nil {
		return nil, err
	}

	blocks := []BlockInterface{}

	for _, blockMap := range blocksMap {
		block, err := mapToBlock(blockMap)

		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}
