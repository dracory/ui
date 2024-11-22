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

func BlocksFromMap(blocks []map[string]any) []BlockInterface {
	blocksMap := []BlockInterface{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, BlockFromMap(block))
	}

	return blocksMap
}

func BlocksToJson(blocks []BlockInterface) (string, error) {
	blocksMap := []blockJsonObject{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, block.ToJsonObject())
	}

	blocksJson, err := json.Marshal(blocksMap)

	return string(blocksJson), err
}

func BlocksToMap(blocks []BlockInterface) []map[string]any {
	blocksMap := []map[string]any{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, block.ToMap())
	}

	return blocksMap
}
