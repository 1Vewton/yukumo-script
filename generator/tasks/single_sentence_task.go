package tasks

import (
	"time"

	"github.com/1Vewton/yukumo-script/data/characters"
)

// SingleSentenceTask defines the task of generating a single sentence
type SingleSentenceTask struct {
	Text      string
	Time      time.Time
	Character *characters.Character
	PhontName *string
}
