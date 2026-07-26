package syncutils

import (
	"sync"
)

// Map stores the examples and also make it memory safe.
type Map struct {
	sync.RWMutex
	examples map[string]string
}

// NewMap creates new Map
func NewMap() *Map {
	return &Map{
		examples: make(map[string]string),
	}
}

// SetKV sets a key-value pair
func (m *Map) SetKV(key string, value string) {
	m.Lock()
	defer m.Unlock()
	m.examples[key] = value
}

// GetValue gets the value of a key-value pair
func (m *Map) GetValue(key string) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	value, exists := m.examples[key]
	return value, exists
}
