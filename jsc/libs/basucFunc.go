package libs

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	BASE_PATH       string            // base path for the Just-Simple-Cloud environment
	STACK_FILE_NAME string = ".stack" // default stack file name
	STACK_FILE      string = ""       // full path to the stack file
)

func Init() {
	// Initialization logic for the libs package

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

}
