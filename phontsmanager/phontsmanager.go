package phontsmanager

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/1Vewton/yukumo-script/utils/syncutils"
)

// PhontNameToFileName creates mapping of phont name and file
var PhontNameToFileName = syncutils.NewMap()

// GetAllPhonts gets all the phonts
func GetAllPhonts(phontsDir string) ([]os.DirEntry, error) {
	return os.ReadDir(phontsDir)
}

// InitializePhontNameToFileName initializes the key-value pair of phont name and phont file
func InitializePhontNameToFileName(phontsDir string) error {
	phonts, err := GetAllPhonts(phontsDir)
	if err != nil {
		return err
	}
	for _, phont := range phonts {
		name := phont.Name()
		extension := filepath.Ext(name)
		if extension == ".phont" {
			phontName := strings.TrimSuffix(
				name,
				extension,
			)
			PhontNameToFileName.SetKV(phontName, name)
		}
	}
	return nil
}
