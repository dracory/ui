package ui

import (
	"reflect"
	"testing"
)

func TestNewFromJSON(t *testing.T) {
	type args struct {
		blockJson string
	}

	tests := []struct {
		name    string
		args    args
		want    *Block
		wantErr bool
	}{
		{
			name: "NewFromJSON",
			args: args{
				blockJson: `{"id":"1","type":"block1","content":"","parameters":{"key":"value"},"children":[]}`,
			},
			want: &Block{
				id:         "1",
				blockType:  "block1",
				content:    "",
				parameters: map[string]string{"key": "value"},
				children:   []*Block{},
			},
			wantErr: false,
		},
		{
			name: "NewFromJSON",
			args: args{
				blockJson: `{"id":"1","type":"block1","content":"","parameters":{"key":"value"},"children":[{"id":"2","type":"block2","content":"","parameters":{"key2":"value2"},"children":[]},{"id":"3","type":"block3","content":"","parameters":{"key3":"value3"},"children":[]}]}`,
			},
			want: &Block{
				id:         "1",
				blockType:  "block1",
				content:    "",
				parameters: map[string]string{"key": "value"},
				children: []*Block{
					{
						id:         "2",
						blockType:  "block2",
						content:    "",
						parameters: map[string]string{"key2": "value2"},
						children:   []*Block{},
					},
					{
						id:         "3",
						blockType:  "block3",
						content:    "",
						parameters: map[string]string{"key3": "value3"},
						children:   []*Block{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromJson(tt.args.blockJson)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
