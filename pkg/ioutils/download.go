package ioutils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
)


// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(url, localFile string) error {

	// Create the file
	out, err := os.Create(localFile)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}


func ReadStreamFile(url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Sprintf("%v", resp.Body)

	return nil
}


// DownloadCommandLinePrograms checks for downloads (if does not yet exist on
// $PATH): bigWigToWig
func DownloadCommandLinePrograms(config *Configuration) error {
	// *********************** bigWigToWig *************************
	p, _ := GetExecutable("bigWigToWig")
	if p == "" {
		err := downloadbigWigToWig(config)
		if err != nil {
			return err
		}
	} else {
		config.bigWigToWig = p
	}
	return nil
}


func GetExecutable(program string) (string, error) {
	path, err := exec.LookPath(program)
	return path, err
}

func downloadbigWigToWig(config *Configuration) error {
	var sourcePath string
	if config.OS == "darwin" {
		sourcePath = "http://hgdownload.cse.ucsc.edu/admin/exe/macOSX.x86_64/bigWigToWig"
	} else {
		sourcePath = "http://hgdownload.cse.ucsc.edu/admin/exe/linux.x86_64/bigWigToWig"
	}

	outPath := path.Join(config.SoftwareDownloadDIR, "bigWigToWig")
	DownloadFile(sourcePath, outPath)
	fmt.Println(outPath)
	err := os.Chmod(outPath, 0777)
	if err != nil {
		return err
	}
	config.bigWigToWig = outPath
	return nil
}
