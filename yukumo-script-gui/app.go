//go:build windows
// +build windows

package main

import (
	"context"
	"fmt"

	"github.com/1Vewton/yukumo-script/example"
	"github.com/1Vewton/yukumo-script/phontsmanager"
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

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	// Initialize utils
	utils.InitializeDirectory(utils.PhontsDir)
	utils.InitializeDirectory(utils.ResultDir)
	utils.InitializeDirectory(utils.WavsDir)
	utils.InitializeDirectory(utils.DatasDir)
	utils.InitializeDirectory(utils.ExampleDir)
	utils.InitializeDirectory(utils.ImagesDir)
	dir, err := phontsmanager.GetAllPhonts(
		utils.PhontsDir,
	)
	if err != nil {
		guiLogger.Error(err.Error())
		panic(err.Error())
	}
	err = example.GenerateExampleWin(
		ctx,
		utils.ExampleDir,
		utils.PhontsDir,
		dir,
	)
	if err != nil {
		guiLogger.Error(err.Error())
		panic(err.Error())
	}
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
