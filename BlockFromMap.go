package ui

import "reflect"

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
