package libs

import (
	"log"
	"os"
	"path/filepath"
)

var (
	BASE_PATH string = ""
)

func Init() {
	// Initialization logic for the libs package

	if os.Getenv("BASE_PATH") != "" {
		p, err := filepath.Abs(os.Getenv("BASE_PATH"))
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

}
