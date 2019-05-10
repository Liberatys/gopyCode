package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var mu sync.Mutex
var file *os.File
var routines sync.WaitGroup

func createNewFile(fileName string) {
	mu.Lock()
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	file = f
	mu.Unlock()
}

func writeToFile(data []byte) {
	mu.Lock()
	defer routines.Done()
	file.Write([]byte(data))
	mu.Unlock()
}

func closeFile() {
	mu.Lock()
	err := file.Close()
	if err != nil {
		panic(err)
	}
	mu.Unlock()
}

func addRoutine() {
	routines.Add(1)
}

func readFileData(filePath string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	dir := filepath.Dir(filePath)
	fileName := "| " + filepath.Base(dir) + "/" + filepath.Base(filePath) + " |"
	fileNameLength := len(fileName)
	fileName = strings.Repeat("-", fileNameLength-3) + "+\n" + fileName
	data := "+ " + fileName + "\n+ " + strings.Repeat("-", fileNameLength-3) + "+\n\n\n" + string(b[:]) + "\n\n\n"
	addRoutine()
	go writeToFile([]byte(data))
}
