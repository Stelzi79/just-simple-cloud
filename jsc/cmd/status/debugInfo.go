package status

import (
	"github.com/Stelzi79/just-simple-cloud/libs"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func DebugInfo(cmd *cobra.Command) {
	if debug, _ := cmd.Flags().GetBool("debug"); debug {
		libs.ColorPrintln(color.FgRed, "🐛 Debug mode is enabled.")
	}

}
