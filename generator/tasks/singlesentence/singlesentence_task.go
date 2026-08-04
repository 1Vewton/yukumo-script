package singlesentence

import (
	"errors"
	"time"

	"github.com/1Vewton/yukumo-script/data/characters"
)

// SingleSentenceTask defines the task of generating a single sentence
type SingleSentenceTask struct {
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
) (*SingleSentenceTask, error) {
	if phontName == nil && character == nil {
		return nil, errors.New(
			"You have to choose at least one of the way to generate the audio",
		)
	}
	return &SingleSentenceTask{
		Text:       text,
		CreateTime: time.Now(),
		EditTime:   time.Now(),
		Character:  character,
		PhontName:  phontName,
	}, nil
}
