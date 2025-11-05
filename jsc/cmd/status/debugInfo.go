package status

import (
	"log"
	"os"

	"github.com/Stelzi79/just-simple-cloud/libs"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func DebugInfo(cmd *cobra.Command) {
	if debug, _ := cmd.Flags().GetBool("debug"); debug {
		log.Println("🐛 Debug mode is enabled.")

		if env, _ := cmd.Flags().GetBool("env"); env {
			// Log all Environment Variables
			log.Println("Environment Variables:")
			for _, env := range os.Environ() {
				log.Println(env)
				color.Unset()
			}
		}

		log.Println("`BASE_PATH` set to:", libs.BASE_PATH)
		log.Println("`STACK_FILE_NAME` set to:", libs.STACK_FILE_NAME)
		log.Println("`STACK_FILE` set to:", libs.STACK_FILE)
		log.Printf("Loaded stack file: %+v", libs.STACK_DATA.PrettyPrint())
	}
}
