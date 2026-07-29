package characters

import (
	"sync"
)

// Characters stores the list of characters
type Characters struct {
	sync.RWMutex
	Data []Character `json:"data"`
}

// NewCharacters creates new Characters
func NewCharacters() *Characters {
	return &Characters{
		Data: []Character{},
	}
}
