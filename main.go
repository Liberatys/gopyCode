package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Gopier struct {
	startFolder string
	outputFile  string
	extensions  []string
	timer       int64
	timed       bool
}

func (gopier *Gopier) checkForDefaults() {
	if gopier.startFolder == "" {
		currentFilePath, err := filepath.Abs("./")
		if err != nil {
			panic("Not able to find current folder")
		}
		gopier.startFolder = currentFilePath
	}
	if gopier.outputFile == "" {
		gopier.outputFile = "output.txt"
	}
	if len(gopier.extensions) == 0 {
		panic("Please set extensions to use")
	}
}

func main() {
	arguments := os.Args[1:]
	setFlags()
	gopier := parseFlags(arguments)
	if gopier.timed {
		gopier.timer = time.Now().UnixNano()
	}
	gopier.checkForDefaults()
	files, err := getFileListOfDirectory(gopier.startFolder)
	if err != nil {
		panic(err)
	}
	var FileExtensions map[string][]string
	FileExtensions = make(map[string][]string)
	createNewFile(gopier.outputFile)
	for extension := range gopier.extensions {
		FileExtensions[gopier.extensions[extension]] = make([]string, 0)
	}
	for i := range files {
		for j := range gopier.extensions {
			if path.Ext(files[i]) == gopier.extensions[j] {
				FileExtensions[gopier.extensions[j]] = append(FileExtensions[gopier.extensions[j]], files[i])
			}
		}
	}
	mapping := make(map[string][]string)
	for key, value := range FileExtensions {
		var newValues = make([]string, len(value))
		for i := range value {
			newValues[i] = filepath.Base(filepath.Dir(value[i])) + "/" + filepath.Base(value[i])
		}
		mapping[key] = newValues
	}
	//display overview map on top of the file
	data, _ := json.MarshalIndent(mapping, "", "    ")
	writing := string(data[:]) + "\n\n\n"
	addRoutine()
	writeToFile([]byte(writing))
	for _, value := range FileExtensions {
		for i := range value {
			readFileData(value[i])
		}
		routines.Wait()
		addRoutine()
		writeToFile([]byte(strings.Repeat("-", 80) + "\n"))
	}
	routines.Wait()
	closeFile()
	if gopier.timed {
		fmt.Println(fmt.Sprintf("Duration: %v ms", (time.Now().UnixNano()-gopier.timer)/1000000))
	}
}
