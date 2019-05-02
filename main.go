package main

import (
	"os"
	"path"
	"path/filepath"
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
	for _, value := range FileExtensions {
		for i := range value {
			readFileData(value[i])
		}
	}
	closeFile()
}
