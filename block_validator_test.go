// block_validator_test.go
package ui

import (
	"errors"
	"testing"
)

func TestBlockValidator_Add(t *testing.T) {
	validator := NewBlockValidator()

	// Test adding a validator
	called := false
	testValidator := func(block BlockInterface) error {
		called = true
		return nil
	}

	validator.Add("test", testValidator)

	// Test that the correct validator is called
	block := NewBlock()
	block.SetType("test")
	_ = validator.Validate(block)

	if !called {
		t.Error("Expected validator to be called")
	}
}

func TestBlockValidator_Validate(t *testing.T) {
	validator := NewBlockValidator()

	// Setup test cases
	tests := []struct {
		name        string
		blockType   string
		setup       func(*BlockValidator)
		wantErr     bool
		errContains string
	}{
		{
			name:      "no validator for type",
			blockType: "unknown",
			setup:     func(*BlockValidator) {},
			wantErr:   false,
		},
		{
			name:      "validator succeeds",
			blockType: "button",
			setup: func(v *BlockValidator) {
				v.Add("button", func(BlockInterface) error { return nil })
			},
			wantErr: false,
		},
		{
			name:      "validator fails",
			blockType: "button",
			setup: func(v *BlockValidator) {
				v.Add("button", func(BlockInterface) error {
					return errors.New("validation failed")
				})
			},
			wantErr:     true,
			errContains: "validation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(validator)

			var block BlockInterface
			if tt.blockType != "" {
				block = NewBlock()
				block.SetType(tt.blockType)
			}

			err := validator.Validate(block)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if tt.errContains != "" && err.Error() != tt.errContains {
					t.Errorf("error %q does not contain %q", err.Error(), tt.errContains)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestBlockValidator_NilBlock(t *testing.T) {
	validator := NewBlockValidator()
	
	// Add a validator that will panic if called
	validator.Add("button", func(BlockInterface) error {
		t.Error("validator should not be called for nil block")
		return nil
	})
	
	// Should not panic and should return nil error
	if err := validator.Validate(nil); err != nil {
		t.Errorf("unexpected error for nil block: %v", err)
	}
}

func TestBlockValidator_Concurrent(t *testing.T) {
	validator := NewBlockValidator()

	// Start multiple goroutines that add validators
	const numGoroutines = 10
	done := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			blockType := string(rune('a' + n)) // Create unique block type for each goroutine
			validator.Add(blockType, func(BlockInterface) error { return nil })
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	// Verify all validators were added
	block := NewBlock()
	for i := 0; i < numGoroutines; i++ {
		blockType := string(rune('a' + i))
		block.SetType(blockType)
		if err := validator.Validate(block); err != nil {
			t.Errorf("validator for type %s not found", blockType)
		}
	}
}
