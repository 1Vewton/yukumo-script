package generatorwin

/*
#cgo LDFLAGS: -L${SRCDIR}/../../lib/win64 -l:AquesTalk2.lib
#include<YukumoGenerator.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// GeneratorWin generates yukumo audio for windows
type GeneratorWin struct {
	speed      int
	phontPath  string
	resultPath string
	text       string
}

// NewGeneratorWin creates new GeneratorWin struct
func NewGeneratorWin(
	speed int,
	phontPath string,
	resultPath string,
	text string,
) *GeneratorWin {
	return &GeneratorWin{
		speed:      speed,
		phontPath:  phontPath,
		resultPath: resultPath,
		text:       text,
	}
}

// generate_wav generates yukumo .wav file on Windows
func (winGenerator *GeneratorWin) GenerateWav() error {
	tmpPhontPath := C.CString(winGenerator.phontPath)
	tmpText := C.CString(winGenerator.text)
	tmpResultPath := C.CString(winGenerator.resultPath)
	result := C.generate_wav(
		tmpPhontPath,
		tmpText,
		tmpResultPath,
		C.int(winGenerator.speed),
	)
	C.free(unsafe.Pointer(tmpPhontPath))
	C.free(unsafe.Pointer(tmpText))
	C.free(unsafe.Pointer(tmpResultPath))
	switch result {
	case 0:
		return nil
	case -1:
		return errors.New("File loading error")
	case -2:
		return errors.New("wav generating error")
	case -3:
		return errors.New("File opening error")
	case -4:
		return errors.New("write incomplete")
	default:
		return fmt.Errorf(
			"Unexpected return %d!",
			result,
		)
	}
}
