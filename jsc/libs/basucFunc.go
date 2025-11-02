package libs

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	BASE_PATH string
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

}
