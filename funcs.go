package ui

import (
	"encoding/json"
	"errors"
)

func MarshalBlocksToJson(blocks []BlockInterface) (string, error) {
	blocksMap := []blockJsonObject{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, block.ToJsonObject())
	}

	blocksJson, err := json.Marshal(blocksMap)

	return string(blocksJson), err
}

func UnmarshalJsonToBlocks(blocksJson string) ([]BlockInterface, error) {
	blocksMap := []map[string]any{}

	err := json.Unmarshal([]byte(blocksJson), &blocksMap)

	if err != nil {
		return nil, err
	}

	blocks := []BlockInterface{}

	for _, blockMap := range blocksMap {
		block, err := ConvertMapToBlock(blockMap)

		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}

func ConvertMapToBlocks(blocks []map[string]any) []BlockInterface {
	blocksMap := []BlockInterface{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, NewBlockFromMap(block))
	}

	return blocksMap
}

func ConvertBlocksToMap(blocks []BlockInterface) []map[string]any {
	blocksMap := []map[string]any{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, block.ToMap())
	}

	return blocksMap
}

// ConvertMapToBlock converts a map to a block
//
// The map must represent a valid block (have parameters like id, and type),
// otherwise an error will be returned
//
// Parameters:
// - blockMap - a map[string]any to convert to a block
//
// Returns:
// - BlockInterface - a block
// - error - if the map[string]any is not a valid block
func ConvertMapToBlock(blockMap map[string]any) (BlockInterface, error) {
	blockMap, err := mapToBlockMap(blockMap)

	if err != nil {
		return nil, err
	}

	return NewBlockFromMap(blockMap), nil
}

// mapToBlockMap converts a map[string]any to a map[string]any
// the map[string]any must be a valid block, otherwise an error
// will be returned
//
// Parameters:
// - blockMap - a map[string]any to convert to a block
//
// Returns:
// - map[string]any - a block
// - error - if the map[string]any is not a valid block
func mapToBlockMap(blockMap map[string]any) (map[string]any, error) {
	idAny, ok := blockMap["id"]

	if !ok {
		return nil, errors.New("id not found")
	}

	typeAny, ok := blockMap["type"]

	if !ok {
		return nil, errors.New("type not found")
	}

	parametersAny, ok := blockMap["parameters"]

	if !ok {
		parametersAny = map[string]any{}
	}

	childrenAny, ok := blockMap["children"]

	if !ok {
		childrenAny = []any{}
	}

	childrenArrayAny := childrenAny.([]any)

	childrenMap := []map[string]any{}
	for _, childAny := range childrenArrayAny {
		childAny := childAny.(map[string]any)
		child, err := mapToBlockMap(childAny)

		if err != nil {
			return nil, err
		}

		childrenMap = append(childrenMap, child)
	}

	parametersMapAny := parametersAny.(map[string]any)
	parametersMap := map[string]string{}

	for k, v := range parametersMapAny {
		parametersMap[k] = v.(string)
	}

	blockMap["id"] = idAny.(string)
	blockMap["type"] = typeAny.(string)
	blockMap["parameters"] = parametersMap
	blockMap["children"] = childrenMap

	return blockMap, nil
}
