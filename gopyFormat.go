package main

import (
	"encoding/json"
	"runtime"
	"time"
)

type GopyForamt struct {
	BaseFolder    string
	FileDelimiter string
	TableView     map[string][]string
	MetaData      MetaData
	FilesFound    int `json:"fileFound"`
}

func assembleGopyFormat(tableView map[string][]string, baseFolder string) GopyForamt {
	gopyFormat := GopyForamt{
		MetaData:      assembleMetaData(),
		FileDelimiter: "--:",
		TableView:     tableView,
		BaseFolder:    baseFolder,
	}
	return gopyFormat
}

func (gopyFormat *GopyForamt) convertToJSON() string {
	data, _ := json.MarshalIndent(gopyFormat, "", "    ")
	return string(data[:])
}

type MetaData struct {
	CreationDate string `json:"creationDate"`
	HostPlatform string `json:"creationOS"`
	Version      string `json:"creationVersion"`
}

func assembleMetaData() MetaData {
	metaData := MetaData{
		CreationDate: time.Now().Format("01-02-2006"),
		HostPlatform: runtime.GOOS,
		Version:      version,
	}
	return metaData
}
