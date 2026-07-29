package characters

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Characters stores the list of characters
type Characters struct {
	sync.RWMutex
	Data   []Character `json:"data"`
	Folder string      `json:"folder"`
	File   string      `json:"file"`
}

// NewCharacters creates new Characters
func NewCharacters() *Characters {
	return &Characters{
		Data: []Character{},
	}
}

// SetTargetFile sets the target file to store the Characters
func (characters *Characters) SetTargetFile(
	folder string,
	file string,
) {
	characters.Folder = folder
	characters.File = file
}

// AddCharacter adds new character to the slice
func (characters *Characters) AddCharacter(character *Character) {
	if character != nil {
		characters.Lock()
		defer characters.Unlock()
		characters.Data = append(characters.Data, *character)
	}
}

// ReadData reads the data inside the file stored
func (characters *Characters) ReadData() error {
	path := fmt.Sprintf(
		"%s/%s",
		characters.Folder,
		characters.File,
	)
	_, errExist := os.Stat(path)
	if errExist != nil {
		if os.IsNotExist(errExist) {
			file, errCreate := os.Create(path)
			if errCreate != nil {
				return errCreate
			}
			defer file.Close()
			jsonData, errJson := json.Marshal(characters)
			if errJson != nil {
				return errJson
			}
			_, errWrite := file.Write(jsonData)
			return errWrite
		} else {
			return errExist
		}
	}
	data, errRead := os.ReadFile(path)
	if errRead != nil {
		return errRead
	}
	return json.Unmarshal(data, characters)
}
