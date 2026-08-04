package win

import (
	"github.com/spf13/cobra"
)

// generateByFile generates wav by file
var generateByFileCMD = &cobra.Command{
	Use:   "generateByFile",
	Short: "Generate yukumo audio through the file",
	Long: `
generateByFile allows you to generate yukumo audio through phont file directly
	`,
}
