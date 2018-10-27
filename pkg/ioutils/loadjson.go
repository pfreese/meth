package ioutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuration struct {
	DataDIR    string
	SoftwareDownloadDIR    string
}

func ParseJSON(jsonF string) Configuration {
	file, _ := os.Open(jsonF)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}


func cleanConfig(config Configuration) {
	// Remove trailing slashes from directories
	config.DataDIR = filepath.Clean(config.DataDIR)
	config.SoftwareDownloadDIR = filepath.Clean(config.SoftwareDownloadDIR)
}


func checkJSON(config Configuration) {
	// Check that the Data & SoftwareDownload directories are valid
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		// path/to/whatever does not exist
	}
}
