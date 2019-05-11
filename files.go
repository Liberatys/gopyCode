package main

import (
	"os"
	"path/filepath"
)

/*
	Create a struct for later use of more work with the file system and the files that are found in the filesystem.
	For the moment just a placeholder for the correct implementation.
*/

type FileListing struct {
	root         string
	filesLocated int
	filesFound   bool
}

func newFileListing(root string) FileListing {
	return FileListing{
		root: root,
	}
}

func (fileListing *FileListing) getFileListOfDirectory() ([]string, error) {
	var files []string
	err := filepath.Walk(fileListing.root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	fileListing.filesLocated = len(files)
	if len(files) > 0 {
		fileListing.filesFound = true
	}
	return files, err
}
