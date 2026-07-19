package utils

import (
	"fmt"
	"os"

	"github.com/1Vewton/yukumo-script/utils/logger"
)

var initializeLogger = logger.NewLogger(
	"Initialization",
	nil,
)

// InitializeDirectory creates a directory
func InitializeDirectory(directoryName string) {
	_, errStat := os.Stat(directoryName)
	if errStat == nil {
		initializeLogger.Info(
			fmt.Sprintf(
				"%s directory alredy exists \n",
				directoryName,
			),
		)
	} else if os.IsNotExist(errStat) {
		err := os.Mkdir(directoryName, 0666)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(errStat.Error())
	}
}
