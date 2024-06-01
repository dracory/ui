package ui

import (
	"reflect"
	"testing"
)

func TestNewFromMap(t *testing.T) {
	args := map[string]interface{}{
		"id":         "1",
		"type":       "block1",
		"content":    "",
		"parameters": map[string]string{"key": "value"},
		"children": []map[string]interface{}{
			{
				"id":         "2",
				"type":       "block2",
				"content":    "",
				"parameters": map[string]string{"key2": "value2"},
				"children":   []map[string]interface{}{},
			},
			{
				"id":         "3",
				"type":       "block3",
				"content":    "",
				"parameters": map[string]string{"key3": "value3"},
				"children":   []map[string]interface{}{},
			},
		},
	}

	tests := []struct {
		name string
		args map[string]interface{}
		want *Block
	}{
		{
			name: "NewFromMap",
			args: args,
			want: &Block{
				id:         "1",
				blockType:  "block1",
				content:    "",
				parameters: map[string]string{"key": "value"},
				children: []BlockInterface{
					&Block{
						id:         "2",
						blockType:  "block2",
						content:    "",
						parameters: map[string]string{"key2": "value2"},
						children:   []BlockInterface{},
					},
					&Block{
						id:         "3",
						blockType:  "block3",
						content:    "",
						parameters: map[string]string{"key3": "value3"},
						children:   []BlockInterface{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromMap(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
