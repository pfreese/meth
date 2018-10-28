package gtf

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/pfreese/meth/pkg/ioutils"
)

type GTFConfig struct {
	DataDIR    string
	GTF	string
}


func ParseGTFjson(jsonF string) GTFConfig {
	file, _ := os.Open(jsonF)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := GTFConfig{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}


func makeDIRs(config GTFConfig) {

	if _, err := os.Stat(config.DataDIR); os.IsNotExist(err) {
		err := os.Mkdir(config.DataDIR, 0744)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
}


func cleanConfig(config *GTFConfig) {
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



func ProcessGTFjson(jsonF string) (GTFConfig, error) {
	config := ParseGTFjson(jsonF)
	fmt.Println(config)
	cleanConfig(&config)
	makeDIRs(config)
	fmt.Println(config)
	return config, nil
}



func DownloadGTF(config *GTFConfig) error {
	basename := "ENCFF824ZKD.gtf"
	encodeDownload := "https://www.encodeproject.org/files/ENCFF824ZKD/@@download/ENCFF824ZKD.gtf.gz"

	localGTF := path.Join(config.DataDIR, basename)
	if _, err := os.Stat(localGTF); os.IsNotExist(err) {
		err := ioutils.DownloadFile(encodeDownload, localGTF + ".gz")
		if err != nil {
			return err
		}
		_, err = ioutils.UnpackGzipFile(localGTF + ".gz", localGTF)
		if err != nil {
			return err
		}
	}
	config.GTF = localGTF
	return nil
}
