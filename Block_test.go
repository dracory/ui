package ui

import (
	"reflect"
	"testing"
)

func TestBlock_ToMap(t *testing.T) {
	block1 := NewBlock()
	block1.SetID("1")
	block1.SetType("block1")
	block1.SetParameter("key", "value")

	block2 := NewBlock()
	block2.SetID("2")
	block2.SetType("block2")
	block2.SetParameter("key2", "value2")

	block3 := NewBlock()
	block3.SetID("3")
	block3.SetType("block3")
	block3.SetParameter("key3", "value3")

	block1.AddChild(block2)
	block1.AddChild(block3)

	tests := []struct {
		name string
		b    BlockInterface
		want map[string]interface{}
	}{
		{
			name: "Block_ToMap",
			b:    block1,
			want: map[string]interface{}{
				"id":         "1",
				"type":       "block1",
				"parameters": map[string]string{"key": "value"},
				"children": []map[string]interface{}{
					{
						"id":         "2",
						"type":       "block2",
						"parameters": map[string]string{"key2": "value2"},
						"children":   []map[string]interface{}{},
					},
					{
						"id":         "3",
						"type":       "block3",
						"parameters": map[string]string{"key3": "value3"},
						"children":   []map[string]interface{}{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Block.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_ToJson(t *testing.T) {
	block1 := NewBlock()
	block1.SetID("1")
	block1.SetType("block1")
	block1.SetParameter("key", "value")

	block2 := NewBlock()
	block2.SetID("2")
	block2.SetType("block2")
	block2.SetParameter("key2", "value2")

	block3 := NewBlock()
	block3.SetID("3")
	block3.SetType("block3")
	block3.SetParameter("key3", "value3")

	block1.AddChild(block2)
	block1.AddChild(block3)

	block4 := NewBlock()
	block4.SetID("4")
	block4.SetType("block4")

	tests := []struct {
		name    string
		b       BlockInterface
		want    string
		wantErr bool
	}{
		{
			name:    "Block_ToJson",
			b:       block4,
			want:    `{"id":"4","type":"block4","content":"","parameters":{},"children":[]}`,
			wantErr: false,
		},
		{
			name:    "Block_ToJson",
			b:       block1,
			want:    `{"id":"1","type":"block1","content":"","parameters":{"key":"value"},"children":[{"id":"2","type":"block2","content":"","parameters":{"key2":"value2"},"children":[]},{"id":"3","type":"block3","content":"","parameters":{"key3":"value3"},"children":[]}]}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.ToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("Block.ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Block.ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_BlockInterfaceToBlock(t *testing.T) {
	type test struct {
		Block
	}

	testF := func(block BlockInterface) string {
		blockTest := &test{}
		blockTest.Block = *(block.(*Block))
		return blockTest.Type()
	}

	b := NewBlock().SetType("TEST")

	if testF(b) != "TEST" {
		t.Error("Type must be TEST, found", testF(b))
	}
}
