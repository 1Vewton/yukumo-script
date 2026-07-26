package example

import (
	"sync"
)

// ExampleMap stores the examples and also make it memory safe.
type ExampleMap struct {
	sync.RWMutex
	examples map[string]string
}

// NewExampleMap creates new ExampleMap
func NewExampleMap() *ExampleMap {
	return &ExampleMap{
		examples: make(map[string]string),
	}
}

// SetKV sets a key-value pair
func (m *ExampleMap) SetKV(key string, value string) {
	m.Lock()
	defer m.Unlock()
	m.examples[key] = value
}

// GetValue gets the value of a key-value pair
func (m *ExampleMap) GetValue(key string) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	value, exists := m.examples[key]
	return value, exists
}
