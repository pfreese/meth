package main

import (
	"fmt"
	"github.com/pfreese/meth/pkg/ioutils"
)

func main() {

	//fileUrl := "https://www.nytimes.com"
	fileUrl := "http://hgdownload.cse.ucsc.edu/admin/exe/macOSX.x86_64/bigWigToWig"

	//err := DownloadFile(fileUrl,"/Users/pfreese/Downloads/nyt.txt")
	err := ioutils.DownloadFile(fileUrl,"/Users/pfreese/Downloads/bigWigToWigGG")
	if err != nil {
		panic(err)
	}

	// Parse the json file
	config := ioutils.ParseJSON("../configs/setup.json")
	fmt.Println(config)
}
