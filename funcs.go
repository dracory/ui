package ui

import "errors"

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
		//return nil, errors.New("parameters not found")
		parametersAny = map[string]any{}
	}

	childrenAny, ok := blockMap["children"]

	if !ok {
		// return nil, errors.New("children not found")
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

func mapToBlock(blockMap map[string]any) (BlockInterface, error) {
	blockMap, err := mapToBlockMap(blockMap)

	if err != nil {
		return nil, err
	}

	return BlockFromMap(blockMap), nil
}
