package generatorwin

// #cgo LDFLAGS: -L${SRCDIR}/../lib/win64 -lAquesTalk2
import "C"

// GeneratorWin generates yukumo audio for windows
type GeneratorWin struct {
	speed       int
	phont_path  string
	result_path string
	text        string
}

// NewGeneratorWin creates new GeneratorWin struct
func NewGeneratorWin(
	speed int,
	phont_path string,
	result_path string,
	text string,
) *GeneratorWin {
	return &GeneratorWin{
		speed:       speed,
		phont_path:  phont_path,
		result_path: result_path,
		text:        text,
	}
}
