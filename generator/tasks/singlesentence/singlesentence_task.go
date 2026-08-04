package singlesentence

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/1Vewton/yukumo-script/data/characters"
	"github.com/google/uuid"
)

// Task defines the task of generating a single sentence
type Task struct {
	ID         string                `json:"id"`
	TaskName   string                `json:"taskName"`
	Text       string                `json:"text"`
	CreateTime time.Time             `json:"createTime"`
	EditTime   time.Time             `json:"editTime"`
	Character  *characters.Character `json:"character"`
	PhontName  *string               `json:"phontName"`
}

// NewSingleSentenceTask creates new single sentence task
func NewSingleSentenceTask(
	text string,
	character *characters.Character,
	phontName *string,
) (*Task, error) {
	if phontName == nil && character == nil {
		return nil, errors.New(
			"You have to choose at least one of the way to generate the audio",
		)
	}
	id := uuid.NewString()
	return &Task{
		ID:         id,
		Text:       text,
		CreateTime: time.Now(),
		EditTime:   time.Now(),
		Character:  character,
		PhontName:  phontName,
	}, nil
}

// NewSingleSentenceTaskFromFile gets single sentence task from file
func NewSingleSentenceTaskFromFile(
	fileName string,
) (*Task, error) {
	var result Task
	data, errRead := os.ReadFile(fileName)
	if errRead != nil {
		return nil, errRead
	}
	errJSON := json.Unmarshal(data, &result)
	if errJSON != nil {
		return nil, errJSON
	}
	return &result, nil
}

// GenerateFileName generates filename of metadata for this task
func (task *Task) GenerateFileName(
	targetDir string,
) string {
	return fmt.Sprintf(
		"%s/%s_%s_%d.json",
		targetDir,
		task.TaskName,
		task.ID,
		task.CreateTime.Unix(),
	)
}

// SaveFileWin saves the file in the target directory in windows
func (task *Task) SaveFileWin(
	targetDir string,
) error {
	// Marshal
	marshalResult, errMarshal := json.Marshal(*task)
	if errMarshal != nil {
		return errMarshal
	}
	// Write file
	fileName := task.GenerateFileName(targetDir)
	errWrite := os.WriteFile(
		fileName,
		marshalResult,
		0644,
	)
	return errWrite
}
