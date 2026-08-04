package singlesentence

import (
	"fmt"
	"os"

	"github.com/1Vewton/yukumo-script/utils/logger"
)

var singleSentenceManagerLogger = logger.NewLogger(
	"SingleSentenceManager",
	nil,
)

// GetAllTasks gets all the tasks
func GetAllTasks(
	targetDir string,
) ([]*SingleSentenceTask, error) {
	result := []*SingleSentenceTask{}
	allFiles, errRead := os.ReadDir(targetDir)
	if errRead != nil {
		return nil, errRead
	}
	for _, file := range allFiles {
		tmpFilePath := fmt.Sprintf(
			"%s/%s",
			targetDir,
			file.Name(),
		)
		data, errGet := NewSingleSentenceTaskFromFile(
			tmpFilePath,
		)
		if errGet != nil {
			singleSentenceManagerLogger.Error(
				fmt.Sprintf(
					"Data in file %s cannot be read because of %s",
					tmpFilePath,
					errGet.Error(),
				),
			)
		} else {
			result = append(result, data)
		}
	}
	return result, nil
}
