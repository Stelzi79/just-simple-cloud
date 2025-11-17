package types

import "os"

var (
	BASE_PATH    string   // base path for the Just-Simple-Cloud environment
	RootBasePath *os.Root // root opened base path

	// Default values for cloud file
	CLOUD_FILE_NAME string    = ".cloud" // default cloud file name
	CLOUD_FILE      string    = ""       // full path to the cloud file
	CLOUD_DATA      CloudFile            // parsed cloud file data

	// Default values for stack file
	STACK_FILE_NAME string    = ".stack" // default stack file name
	STACK_FILE      string    = ""       // full path to the stack file
	STACK_DATA      StackFile            // parsed stack file data

)
