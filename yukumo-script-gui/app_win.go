//go:build windows
// +build windows

package main

import (
	"context"

	"github.com/1Vewton/yukumo-script/example"
	"github.com/1Vewton/yukumo-script/phontsmanager"
	"github.com/1Vewton/yukumo-script/utils"
)

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
