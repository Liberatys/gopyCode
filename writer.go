package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	//wrap the filename into a good looking and easy findable box that can be searched for in a better way.
	wrappedFilename := wrapFileName(filePath)
	data := wrappedFilename + string(b[:]) + "\n\n"
	addRoutine()
	go writeToFile([]byte(data))
}

func wrapFileName(filename string) string {
	firstLine := "+ " + strings.Repeat("-", len(filename)) + " +\n"
	middleLine := "| " + filename + " |\n"
	return firstLine + middleLine + firstLine + "\n\n"
}
