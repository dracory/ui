package ui

func NewFromMap(m map[string]interface{}) *Block {
	id := ""

	if idMap, ok := m["id"].(string); ok {
		id = idMap
	}

	blockType := ""

	if blockTypeMap, ok := m["type"].(string); ok {
		blockType = blockTypeMap
	}

	content := ""

	if contentMap, ok := m["content"].(string); ok {
		content = contentMap
	}

	parameters := map[string]string{}

	if parametersMap, ok := m["parameters"].(map[string]string); ok {
		for k, v := range parametersMap {
			parameters[k] = v
		}
	}

	children := []*Block{}

	if childrenAny, ok := m["children"]; ok {
		childrenMap := childrenAny.([]map[string]interface{})
		for _, c := range childrenMap {
			child := NewFromMap(c)
			if child == nil {
				continue
			}
			children = append(children, child)
		}
	}

	block := NewBlock()
	block.SetID(id)
	block.SetType(blockType)
	block.SetContent(content)
	block.SetParameters(parameters)
	block.SetChildren(children)
	return block
}
