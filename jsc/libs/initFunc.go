package libs

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Stelzi79/just-simple-cloud/types"
)

var (
	BASE_PATH string // base path for the Just-Simple-Cloud environment

	// Default values for cloud file
	CLOUD_FILE_NAME string          = ".cloud" // default cloud file name
	CLOUD_FILE      string          = ""       // full path to the cloud file
	CLOUD_DATA      types.CloudFile            // parsed cloud file data

	// Default values for stack file
	STACK_FILE_NAME string          = ".stack" // default stack file name
	STACK_FILE      string          = ""       // full path to the stack file
	STACK_DATA      types.StackFile            // parsed stack file data

)

// Initialization logic for the jsc
func Init() {
	if os.Getenv("BASE_PATH") != "" {
		key := os.Getenv("BASE_PATH")
		// remove potential quotes
		key = strings.ReplaceAll(key, "\"", "")
		// remove trailing whitespace
		key = strings.TrimSpace(key)
		// get absolute path
		p, err := filepath.Abs(key)
		if err != nil {
			log.Fatalf("Error getting absolute path %s: %v", os.Getenv("BASE_PATH"), err)
		}
		BASE_PATH = p
	} else {
		// get current working directory
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		BASE_PATH = cwd
	}
	// Check if BASE_PATH exists
	if s, err := os.Stat(BASE_PATH); os.IsNotExist(err) || !s.IsDir() {
		log.Fatalf("BASE_PATH does not exist: '%s'", BASE_PATH)
	}

	// Read Cloud File and validate
	// get CLOUD_FILE_NAME from env variable if set
	if os.Getenv("CLOUD_FILE_NAME") != "" {
		CLOUD_FILE_NAME = os.Getenv("CLOUD_FILE_NAME")
		// remove potential quotes
		CLOUD_FILE_NAME = strings.ReplaceAll(CLOUD_FILE_NAME, "\"", "")
		// remove trailing whitespace
		CLOUD_FILE_NAME = strings.TrimSpace(CLOUD_FILE_NAME)
	}

	// check if .cloud file exists in BASE_PATH
	cloudFilePath := filepath.Join(BASE_PATH, CLOUD_FILE_NAME)
	if s, err := os.Stat(cloudFilePath); os.IsNotExist(err) || s.IsDir() {
		log.Fatalf("'%s' file does not exist in BASE_PATH: '%s'", CLOUD_FILE_NAME, cloudFilePath)
	}
	CLOUD_FILE = cloudFilePath

	// read and validate .cloud file
	if err := readAndValidateCloudFile(); err != nil {
		log.Fatalf("Error reading .cloud file '%s': %v", CLOUD_FILE, err)
	}

	// Read Stack File and validate
	// get STACK_FILE_NAME from env variable if set
	if os.Getenv("STACK_FILE_NAME") != "" {
		STACK_FILE_NAME = os.Getenv("STACK_FILE_NAME")
		// remove potential quotes
		STACK_FILE_NAME = strings.ReplaceAll(STACK_FILE_NAME, "\"", "")
		// remove trailing whitespace
		STACK_FILE_NAME = strings.TrimSpace(STACK_FILE_NAME)
	}

	// check if .stack file exists in BASE_PATH
	stackFilePath := filepath.Join(BASE_PATH, STACK_FILE_NAME)
	if s, err := os.Stat(stackFilePath); os.IsNotExist(err) || s.IsDir() {
		log.Fatalf("'%s' file does not exist in BASE_PATH: '%s'", STACK_FILE_NAME, stackFilePath)
	}
	STACK_FILE = stackFilePath

	// read and validate .stack file
	if err := readAndValidateStackFile(); err != nil {
		log.Fatalf("Error reading .stack file '%s': %v", STACK_FILE, err)
	}
}

func readAndValidateCloudFile() error {
	// panic("unimplemented")
	return nil
}

func readAndValidateStackFile() error {
	stack, err := types.NewStackFile(STACK_FILE)
	if err != nil {
		log.Panicf("Error loading stack file '%s': %v", STACK_FILE, err)
	}

	// Validate that the stack file is of type .rootstack
	// Later we can have other types like substacks, etc.
	if !stack.IsRootStack() {
		log.Fatalf("Invalid stack file type in '%s': expected '.rootstack', got '%s'", STACK_FILE, stack.Type)
	}

	// log.Printf("Loaded stack file: %+v", stack.PrettyPrint())
	STACK_DATA = *stack
	return nil
}
