package ioutils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
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
