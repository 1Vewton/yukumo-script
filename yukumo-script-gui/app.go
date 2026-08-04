package main

import (
	"context"
	"fmt"

	"github.com/1Vewton/yukumo-script/utils"
	"github.com/1Vewton/yukumo-script/utils/logger"
	"github.com/magiconair/properties"
)

var guiLogger = logger.NewLogger(
	"GUI",
	nil,
)

// App struct
type App struct {
	ctx      context.Context
	engTexts *properties.Properties
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Load file from English
	pEn, errFile := properties.LoadFile(
		utils.EnglishTexts,
		properties.UTF8,
	)
	if errFile != nil {
		guiLogger.Error(
			fmt.Sprintf(
				"Failed to load English texts due to %s",
				errFile.Error(),
			),
		)
		panic(errFile.Error())
	}
	return &App{
		engTexts: pEn,
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
