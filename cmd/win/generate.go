package win

import (
	"fmt"

	"github.com/1Vewton/yukumo-script/phontsmanager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// SingleSentenceTaskNameByFile defines the name of the task
var SingleSentenceTaskNameByFile string

// generateByFile generates wav by file
var generateByFileCMD = &cobra.Command{
	Use:   "generateByFile",
	Short: "Generate yukumo audio through the file",
	Long: `
generateByFile allows you to generate yukumo audio through phont file directly
- This is completed through 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define the format of the texts
		title := color.New(color.FgGreen).Add(color.Bold)
		text := color.New(color.Italic)
		// Print info
		title.Println("Here are the available phonts:")
		for _, phontName := range phontsmanager.PhontNameToFileName.GetAllKeys() {
			text.Println(phontName)
		}
		title.Println("Input the name of the phont you want to use to generate audio:")
		// Input
		var phontName string
		fmt.Scan(&phontName)
	},
}
