/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Stelzi79/just-simple-cloud/cmd/status"
	"github.com/Stelzi79/just-simple-cloud/libs"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the current status of the Just-Simple-Cloud environment",
	Long: `Shows the current status of the Just-Simple-Cloud environment, including
information about running services, resource usage, updates, and any active alerts.
`,
	Run: func(cmd *cobra.Command, args []string) {
		libs.Init()
		status.DebugInfo(cmd)

		libs.ColorPrintln(color.FgCyan, "🪄 Status of your ☁️ environment:")
		libs.ColorPrintln(color.FgGreen, "✅ All systems operational.")
		libs.ColorPrintln(color.FgGreen, "🟢 x services running.")
		libs.ColorPrintln(color.FgYellow, "💾 Resource usage: xx% CPU, xx% Memory.")
		libs.ColorPrintln(color.FgHiGreen, "🔔 No active alerts.")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Flag to show debug information
	statusCmd.Flags().BoolP("debug", "d", false, "Show debug information")
	// Flag to print Environment Variables in debug mode
	statusCmd.Flags().BoolP("env", "e", false, "Print Environment Variables in debug mode")
}
