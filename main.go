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
	version     string
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
		fmt.Println(fmt.Sprintf("Missing extension flag or extensions \n"))
		getGopyCodeManual("")
	}
	gopier.version = version
}

const (
	version = "v0.0.1"
)

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
	var assemledFiles map[string][]string
	assemledFiles = make(map[string][]string)
	createNewFile(gopier.outputFile)
	for i := range files {
		fileExtension := path.Ext(files[i])
		if gopier.isInExtensions(fileExtension) {
			if _, ok := assemledFiles[fileExtension]; ok == false {
				assemledFiles[fileExtension] = make([]string, 0)
			}
			filename, _ := filepath.Rel(gopier.startFolder, files[i])
			assemledFiles[fileExtension] = append(assemledFiles[fileExtension], filename)
		}
	}
	gopyFormat := assembleGopyFormat(assemledFiles, filepath.Base(gopier.startFolder))
	data, _ := json.MarshalIndent(gopyFormat, "", "    ")
	writing := string(data[:]) + "\n\n\n"
	addRoutine()
	writeToFile([]byte(writing))
	delimiter := strings.Repeat("-", 80)
	for _, value := range assemledFiles {
		for i := range value {
			readFileData(value[i])
		}
		routines.Wait()
		addRoutine()
		writeToFile([]byte(delimiter + "\n"))
	}
	routines.Wait()
	closeFile()
	if gopier.timed {
		fmt.Println(fmt.Sprintf("Duration: %v ms", (time.Now().UnixNano()-gopier.timer)/1000000))
	}
}

func (gopier *Gopier) isInExtensions(extensions string) bool {
	if gopier.extensions[0] == "." {
		return true
	}
	for i := range gopier.extensions {
		if extensions == gopier.extensions[i] {
			return true
		}
	}
	return false
}
