package brainmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Configuration struct {
	DataDIR    string
}

func parseJSON(jsonF string) (Configuration, error) {
	configuration := Configuration{}
	file, err := os.Open(jsonF)
	defer file.Close()
	if err != nil {
		return configuration, err
}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration, nil
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

	// Remove trailing slashes from any directories
	config.DataDIR = filepath.Clean(config.DataDIR)
}

func checkJSON(config Configuration) error {
	// Check that the Data & SoftwareDownload directories are valid
	if _, err := os.Stat(config.DataDIR); os.IsNotExist(err) {
		err := errors.New(fmt.Sprintf("DataDIR %v does not exist", config.DataDIR))
		return err
	}
	return nil
}

func ProcessJSON(jsonF string) (Configuration, error) {
	config, err := parseJSON(jsonF)
	if err != nil {
		return Configuration{}, err
	}
	fmt.Println(config)
	cleanConfig(&config)

	err = checkJSON(config)
	if err != nil {
		return Configuration{}, err
	}
	fmt.Println(config)
	return config, nil
}
