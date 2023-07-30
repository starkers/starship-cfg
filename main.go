package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/miracl/conflate"

	"github.com/adrg/xdg"
)

func main() {
	// log.Println("CFG", xdg.ConfigHome)

	starshipConfigFile, err := xdg.ConfigFile("starship.toml")
	if err != nil {
		log.Fatal(err)
	}

	inputFilePath, err := xdg.SearchConfigFile("starship.d/")
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) < 1 {
		log.Printf("# no files found in %s\n", inputFilePath)
		os.Exit(1)
	}
	var fileNames []string
	for _, f := range files {
		n := f.Name()
		if strings.Contains(n, ".toml") {
			fullPath := fmt.Sprintf("%s/%s", inputFilePath, n)
			fileNames = append(fileNames, fullPath)
		}
	}

	// define the unmarshallers for the given file extensions, blank extension is the global unmarshaller
	conflate.Unmarshallers = conflate.UnmarshallerMap{
		".toml": conflate.UnmarshallerFuncs{conflate.TOMLUnmarshal},
	}

	// merge multiple config files
	c, err := conflate.FromFiles(fileNames...)
	if err != nil {
		log.Fatal(err)
		return
	}

	// output merged data as toml
	toml, err := c.MarshalTOML()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = os.WriteFile(starshipConfigFile, toml, 0644)
	// log.Printf("# from %d files, wrote: %s\n", len(fileNames), starshipConfigFile)
	if err != nil {
		log.Fatal(err)
	}
}
