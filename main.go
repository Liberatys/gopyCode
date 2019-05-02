package main

import (
	"os"
	"path"
	"path/filepath"
)

func main() {
	currentFilePath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	files, err := getFileListOfDirectory(currentFilePath)
	if err != nil {
		panic(err)
	}
	createNewFile(os.Args[2])
	endings := os.Args[1]
	if endings == "" {
		panic("error")
	}
	var element []string
	for i := range files {
		if path.Ext(files[i]) == endings {
			element = append(element, files[i])
		}
	}
	for i := range element {
		readFileData(element[i])
	}
}
