package win

import (
	"github.com/1Vewton/yukumo-script/utils"
	"github.com/1Vewton/yukumo-script/utils/logger"
	"github.com/spf13/cobra"
)

var cmdLogger = logger.NewLogger("CMD", nil)

// rootCMD defines the root command
var rootCMD = &cobra.Command{
	Use:   "yukumo",
	Short: "yukumo is a program that can generate yukumo audio",
	Long: `
Yukumo is a simple and flexible program that can generate yukumo audio without the need for network connection. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CLIWelcome()
	},
}

func init() {
	// Add flags
	generateByFileCMD.Flags().StringVarP(
		&SingleSentenceTaskNameByFile,
		"name",
		"n",
		"task",
		"Sentence1",
	)
	generateByFileCMD.MarkFlagRequired("name")
	// Add subcommands
	rootCMD.AddCommand(
		showAvailablePhontsCMD,
		playExampleCMD,
		generateByFileCMD,
	)
}

// Execute executes the command
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		panic(err.Error())
	}
}
