package ioutils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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


func GetExecutable(program string) (string, error) {
	path, err := exec.LookPath(program)
	fmt.Println(path)
	return path, err
}

func downloadbigWigToWig(config Configuration) {
	path := "a"
	fmt.Println(path)
}
