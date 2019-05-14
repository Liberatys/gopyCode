package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/*
	Gopier stores and works with the information that is needed for the application to run in a good and easy to maintain way.
	Can later be appended with more information and fucntions, to expand the application and the functionality of gopyCode.
*/
type Gopier struct {
	startFolder string
	outputFile  string
	extensions  []string
	timer       int64
	timed       bool
	version     string
}

// check if flags are set, if not set the default values for variables we can set default values.
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
	fileListing := newFileListing(gopier.startFolder)
	files, err := fileListing.getFileListOfDirectory()
	if fileListing.filesFound == false {
		fmt.Println(fmt.Sprintf("Was not able to find any files in the given start folder: %v", gopier.startFolder))
	}
	if err != nil {
		panic(err)
	}
	var assemledFiles map[string][]string
	assemledFiles = make(map[string][]string)
	createNewFile(gopier.outputFile)
	fileCounter := 0
	/*
		Iterate over all found files and check if the extensions is included in the flag configuration.
	*/
	for i := range files {
		fileExtension := path.Ext(files[i])
		if gopier.isInExtensions(fileExtension) {
			if _, ok := assemledFiles[fileExtension]; ok == false {
				assemledFiles[fileExtension] = make([]string, 0)
			}
			fileCounter++
			filename, _ := filepath.Rel(gopier.startFolder, files[i])
			assemledFiles[fileExtension] = append(assemledFiles[fileExtension], filename)
		}
	}
	/*
		simplify the meta data of the file with a json format, so that it can be parsed back very fast and easy.
	*/
	gopyFormat := assembleGopyFormat(assemledFiles, filepath.Base(gopier.startFolder))
	gopyFormat.FilesFound = fileCounter
	writing := gopyFormat.convertToJSON() + "\n\n\n"
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
