package win

import (
	"github.com/1Vewton/yukumo-script/example"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// showAvailablePhontsCMD shows all the available phonts
var showAvailablePhontsCMD = &cobra.Command{
	Use:   "showAvailablePhonts",
	Short: "showAvailablePhonts shows the name of all available phonts",
	Long: `
ShowAvailablePhonts shows all the available phonts that can be used to generate audio
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title := color.New(color.FgGreen).Add(color.Bold)
		title.Println("Here are the available phonts:")
		text := color.New(color.Italic)
		for _, phontName := range example.GetAllExampleFont() {
			text.Println(phontName)
		}
	},
}
