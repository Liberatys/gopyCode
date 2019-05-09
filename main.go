package main

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 2 {
		panic("Please provide output filename and extensions to search for, like\ngopyCode output.txt .java")
	}
	currentFilePath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	files, err := getFileListOfDirectory(currentFilePath)
	if err != nil {
		panic(err)
	}
	var FileExtensions map[string][]string
	FileExtensions = make(map[string][]string)
	createNewFile(arguments[0])
	arguments = arguments[1:]
	//var element []string
	for i := range files {
		for j := range arguments {
			if path.Ext(files[i]) == arguments[j] {
				if _, ok := FileExtensions[arguments[j]]; ok == false {
					FileExtensions[arguments[j]] = make([]string, 0)
				}
				FileExtensions[arguments[j]] = append(FileExtensions[arguments[j]], files[i])
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
	data, _ := json.MarshalIndent(mapping, "", "    ")
	writing := string(data[:]) + "\n\n\n"
	writeToFile([]byte(writing))
	for _, value := range FileExtensions {
		for i := range value {
			readFileData(value[i])
		}
		writeToFile([]byte(strings.Repeat("-", 80)))
	}
	closeFile()
}
