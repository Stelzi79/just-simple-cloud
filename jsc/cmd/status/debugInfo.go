package status

import (
	"log"
	"os"

	"github.com/Stelzi79/just-simple-cloud/types"
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

		log.Println("`BASE_PATH` set to:", types.BASE_PATH)
    

		// Log Cloud File information
		log.Println("`CLOUD_FILE_NAME` set to:", types.CLOUD_FILE_NAME)
		log.Println("`CLOUD_FILE` set to:", types.CLOUD_FILE)
		log.Printf("Loaded cloud file: %+v", types.CLOUD_DATA.PrettyPrint())

		// Log Stack File information
		log.Println("`STACK_FILE_NAME` set to:", types.STACK_FILE_NAME)
		log.Println("`STACK_FILE` set to:", types.STACK_FILE)
		log.Printf("Loaded stack file: %+v", types.STACK_DATA.PrettyPrint())
	}
}
