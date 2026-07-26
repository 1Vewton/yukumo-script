package example

import (
	"fmt"

	"github.com/1Vewton/yukumo-script/utils/audio"
)

// GetAllExampleFont gets the font name of all available phonts
func GetAllExampleFont() []string {
	return examplesMap.GetAllKeys()
}

// PlayExample plays example of a phont
func PlayExample(phontName string) error {
	fileName, exists := examplesMap.GetValue(phontName)
	if !exists {
		return fmt.Errorf(
			"Example for %s does not exists",
			phontName,
		)
	}
	return audio.PlayWAV(fileName)
}
