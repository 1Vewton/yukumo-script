package characters

import (
	"sync"
)

// Characters stores the list of characters
type Characters struct {
	sync.RWMutex
	data []Character
}
