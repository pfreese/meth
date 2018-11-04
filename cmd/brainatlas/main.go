package main

import (
	"fmt"
	"github.com/pfreese/meth/pkg/brainmap"
)

func parseJSON() brainmap.Configuration {
	// Parse the json file
	config, err := brainmap.ProcessJSON("../../configs/brainatlas.json")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(config)
	return config
}

func main() {
	config := parseJSON()

	brainmap.ProcesshVISpConfig(config)
}
