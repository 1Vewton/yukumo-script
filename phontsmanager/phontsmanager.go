package phontsmanager

import (
	"os"

	"github.com/1Vewton/yukumo-script/utils/syncutils"
)

// PhontNameToFile creates mapping of phont name and file
var PhontNameToFileName = syncutils.NewMap()

// GetAllPhonts gets all the phonts
func GetAllPhonts(phontsDir string) ([]os.DirEntry, error) {
	return os.ReadDir(phontsDir)
}
