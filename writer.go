package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var mu sync.Mutex
var file *os.File

func createNewFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	file = f
}

func writeToFile(data []byte) {
	mu.Lock()
	defer mu.Unlock()
	file.Write([]byte(data))
}

func closeFile() {
	mu.Lock()
	defer mu.Unlock()
	file.Close()
}

func readFileData(filePath string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	data := filepath.Base(filePath) + "\n\n\n" + string(b[:]) + "\n\n\n"
	writeToFile([]byte(data))
}
