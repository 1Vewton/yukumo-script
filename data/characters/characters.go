package characters

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

// Characters stores the list of characters
type Characters struct {
	sync.RWMutex
	Data   map[string]*Character `json:"data"`
	Folder string                `json:"folder"`
	File   string                `json:"file"`
}

// NewCharacters creates new Characters
func NewCharacters() *Characters {
	return &Characters{
		Data: make(map[string]*Character),
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
func (characters *Characters) AddCharacter(
	characterID string,
	character *Character,
) error {
	if character != nil {
		characters.Lock()
		defer characters.Unlock()
		_, exists := characters.Data[characterID]
		if exists {
			return fmt.Errorf(
				"Character with Character ID %s already exists",
				characterID,
			)
		}
		characters.Data[characterID] = character
		return nil
	}
	return errors.New(
		"This characterID already exists",
	)
}

// saveTo saves to the target file
func (characters *Characters) saveTo(target string) error {
	jsonData, errJSON := json.Marshal(characters)
	if errJSON != nil {
		return errJSON
	}
	errWrite := os.WriteFile(
		target,
		jsonData,
		0666,
	)
	return errWrite
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
			return characters.saveTo(path)
		}
		return errExist
	}
	data, errRead := os.ReadFile(path)
	if errRead != nil {
		return errRead
	}
	return json.Unmarshal(data, characters)
}

// SaveData saves the data to the file
func (characters *Characters) SaveData() error {
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
			return characters.saveTo(path)
		}
		return errExist
	}
	return characters.saveTo(path)
}

// GetData gets the slice of characters
func (characters *Characters) GetData() map[string]*Character {
	characters.RLock()
	defer characters.RUnlock()
	return characters.Data
}
