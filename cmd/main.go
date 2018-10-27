package main

import (
	"fmt"
	"github.com/pfreese/meth/pkg/ioutils"
	"runtime"
)

func main() {

	/*
	fileUrl := "http://hgdownload.cse.ucsc.edu/admin/exe/macOSX.x86_64/bigWigToWig"

	err := ioutils.DownloadFile(fileUrl,"/Users/pfreese/Downloads/bigWigToWigGG")
	if err != nil {
		panic(err)
	}
	*/

	// Parse the json file
	config, err := ioutils.ProcessJSON("../configs/setup.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(config)

	ioutils.GetExecutable("bigWigToWig")
}
