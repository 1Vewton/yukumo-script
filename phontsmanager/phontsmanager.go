package phontsmanager

import (
	"os"
)

// GetAllPhonts gets all the phonts
func GetAllPhonts(phontsDir string) ([]os.DirEntry, error) {
	return os.ReadDir(phontsDir)
}
