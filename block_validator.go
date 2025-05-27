// block_validator.go
package ui

import "sync"

// Validator validates a block
type Validator func(block BlockInterface) error

// BlockValidator is a thread-safe registry of block validators
type BlockValidator struct {
	mu         sync.RWMutex
	validators map[string]Validator
}

// NewBlockValidator creates a new BlockValidator
func NewBlockValidator() *BlockValidator {
	return &BlockValidator{
		validators: make(map[string]Validator),
	}
}

// Add registers a validator for a block type
func (v *BlockValidator) Add(blockType string, validator Validator) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.validators[blockType] = validator
}

// Validate validates a block using its registered validator
func (v *BlockValidator) Validate(block BlockInterface) error {
	if block == nil {
		return nil
	}

	v.mu.RLock()
	validator, exists := v.validators[block.Type()]
	v.mu.RUnlock()

	if !exists {
		return nil // No validator for this type
	}

	return validator(block)
}
