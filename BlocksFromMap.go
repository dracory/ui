package ui

func BlocksFromMap(blocks []map[string]any) []BlockInterface {
	blocksMap := []BlockInterface{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, BlockFromMap(block))
	}

	return blocksMap
}

func BlocksToMap(blocks []BlockInterface) []map[string]any {
	blocksMap := []map[string]any{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, block.ToMap())
	}

	return blocksMap
}
