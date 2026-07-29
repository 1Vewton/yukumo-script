//go:build windows
// +build windows

package main

import (
	"context"

	"github.com/1Vewton/yukumo-script/cmd/win"
	"github.com/1Vewton/yukumo-script/example"
	"github.com/1Vewton/yukumo-script/phontsmanager"
	"github.com/1Vewton/yukumo-script/utils"
)

// Initialize directories for storing data
func init() {
	utils.InitializeDirectory(utils.PhontsDir)
	utils.InitializeDirectory(utils.ResultDir)
	utils.InitializeDirectory(utils.WavsDir)
	utils.InitializeDirectory(utils.DatasDir)
	utils.InitializeDirectory(utils.ExampleDir)
	utils.InitializeDirectory(utils.ImagesDir)
	dir, err := phontsmanager.GetAllPhonts(utils.PhontsDir)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	err = example.GenerateExampleWin(
		ctx,
		utils.ExampleDir,
		utils.PhontsDir,
		dir,
	)
	if err != nil {
		panic(err)
	}
}

// Main process
func main() {
	win.Execute()
}
