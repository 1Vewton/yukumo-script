package main

import (
	"context"

	"github.com/1Vewton/yukumo-script/example"
	"github.com/1Vewton/yukumo-script/phontsmanager"
	"github.com/1Vewton/yukumo-script/utils"
)

// Initialize directories for storing data
func init() {
	utils.InitializeDirectory("phonts")
	utils.InitializeDirectory("result")
	utils.InitializeDirectory("wav")
	utils.InitializeDirectory("datas")
	utils.InitializeDirectory("examples")
	ctx := context.Background()
	dir, err := phontsmanager.GetAllPhonts("phonts")
	if err != nil {
		panic(err)
	}
	err = example.GenerateExampleWin(
		"examples",
		"phonts",
		ctx,
		dir,
	)
	if err != nil {
		panic(err)
	}
}

// Main process
func main() {

}
