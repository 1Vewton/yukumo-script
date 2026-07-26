package win

import (
	"github.com/spf13/cobra"
)

// rootCMD defines the root command
var rootCMD = &cobra.Command{
	Use:   "yukumo",
	Short: "yukumo is a program that can generate yukumo audio",
	Long: `
Yukumo is a simple and flexible program that can generate yukumo audio without the need for network connection. 
	`,
}

func init() {
	rootCMD.AddCommand(showAvailablePhontsCMD)
}

// Execute executes the command
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		panic(err.Error())
	}
}
