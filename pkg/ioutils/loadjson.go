package ioutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

type Configuration struct {
	DataDIR    string
	SoftwareDownloadDIR    string
	OS 	string
	bigWigToWig string
}

func ParseMeJSON(jsonF string) Configuration {
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

func cleanConfig(config *Configuration) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	// Check if any config directory begins with ~
	if config.DataDIR == "~" {
		config.DataDIR = dir
	} else if strings.HasPrefix(config.DataDIR, "~/") {
		config.DataDIR = filepath.Join(dir, config.DataDIR[2:])
	}
	if config.SoftwareDownloadDIR == "~" {
		config.SoftwareDownloadDIR = dir
	} else if strings.HasPrefix(config.SoftwareDownloadDIR, "~/") {
		config.SoftwareDownloadDIR = filepath.Join(dir, config.SoftwareDownloadDIR[2:])
	}

	// Remove trailing slashes from any directories
	config.DataDIR = filepath.Clean(config.DataDIR)
	config.SoftwareDownloadDIR = filepath.Clean(config.SoftwareDownloadDIR)
}

func checkJSON(config Configuration) error {
	// Check that the Data & SoftwareDownload directories are valid
	if _, err := os.Stat(config.DataDIR); os.IsNotExist(err) {
		err := errors.New(fmt.Sprintf("DataDIR %v does not exist", config.DataDIR))
		return err
	}
	if _, err := os.Stat(config.SoftwareDownloadDIR); os.IsNotExist(err) {
		err := errors.New(fmt.Sprintf("SoftwareDownloadDIR %v does not exist", config.SoftwareDownloadDIR))
		return err
	}
	return nil
}

func ProcessJSON(jsonF string) (Configuration, error) {
	config := ParseMeJSON(jsonF)
	fmt.Println(config)
	cleanConfig(&config)

	err := checkJSON(config)
	if err != nil {
		return Configuration{}, err
	}
	err = addSystemToConfig(&config)
	if err != nil {
		return Configuration{}, err
	}
	fmt.Println(config)
	return config, nil
}

// addSystemToConfig checks that the operating system is either linux or
// mac, and stores it for downstream tool downloads.
func addSystemToConfig(config *Configuration) error {
	operatingSys := runtime.GOOS
	if operatingSys != "darwin" && operatingSys != "linux" {
		err := errors.New("Package only works on OS darwin (Mac) or linux")
		return err
	}
	config.OS = operatingSys
	return nil
}