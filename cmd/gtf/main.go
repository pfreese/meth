package main

import (
	"fmt"

	"github.com/pfreese/meth/pkg/gtf"
	"github.com/pfreese/meth/pkg/ioutils"
)

func parseJSON() {
	// Parse the json file
	config, err := ioutils.ProcessJSON("../configs/setup.json")
	if err != nil {
		fmt.Print(err.Error())
	}

	// Get paths for
	ioutils.DownloadCommandLinePrograms(&config)

	fmt.Println(config)
}


func parseGTF() {
	config, err := gtf.ProcessGTFjson("../configs/hg38.gtf.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	err = gtf.DownloadGTF(&config)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = gtf.ParseGTF(config.GTF)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func readStream() {
	url := "ftp://ftp.ebi.ac.uk/pub/databases/gencode/Gencode_human/release_29/gencode.v29.tRNAs.gff3.gz"
	err := ioutils.ReadStreamFile(url)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func main() {

	//parseJSON()
	//parseGTF()
	readStream()
}
