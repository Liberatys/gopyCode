package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Flag struct {
	name             string
	flags            []string
	requiresTrailing bool
	description      string
	function         func(...string)
}

var flags = []Flag{}

func addDefaultFlags() {
	helpFlag := Flag{
		name:             "help",
		flags:            []string{"-h", "--h", "-help"},
		requiresTrailing: false,
		description:      "Display usage of gopyCode",
		function:         getGopyCodeManual,
	}
	flags = append(flags, helpFlag)
}

func addFlags() {
	flags = append(flags, Flag{
		name:             "output",
		flags:            []string{"-o", "--o", "-out"},
		requiresTrailing: true,
		description: "Define the ouput file for gopyCode like\n		 -> -o output.txt\n",
		function: setOutPutFile,
	})
	flags = append(flags, Flag{
		name:             "start",
		flags:            []string{"-s", "--s", "-src"},
		requiresTrailing: true,
		description: "Set the starting folder for gopyCode like\n		 -> -src .\n",
		function: setStartingPoint,
	})

	flags = append(flags, Flag{
		name:             "extensions",
		flags:            []string{"-ex", "-e", "--e"},
		requiresTrailing: true,
		description: "Define the extennsions, gopyCode should look for like \n		-> -ex .java .go\n",
		function: setExtensions,
	})
}

func setExtensions(extensions ...string) {
	newGopier.extensions = extensions
}

func setStartingPoint(startDir ...string) {
	folder, err := filepath.Abs(startDir[0])
	if err != nil {
		panic(err.Error())
	}
	newGopier.startFolder = folder
}

func setOutPutFile(filename ...string) {
	newGopier.outputFile = filename[0]
}

func getGopyCodeManual(arg ...string) {
	var manual string
	for i := range flags {
		if flags[i].name != "help" {
			manual += fmt.Sprintf("*  %v	%v\n	 %v\n", flags[i].flags, flags[i].name, flags[i].description)
		}
	}
	fmt.Println("\n" + manual)
}

func setFlags() {
	addDefaultFlags()
	addFlags()
}

var newGopier Gopier

func parseFlags(arguments []string) Gopier {
	newGopier = Gopier{}
	var currentFlag Flag
	var currentArguments []string
	for argument := range arguments {
		if strings.HasPrefix(arguments[argument], "-") || strings.HasPrefix(arguments[argument], "--") {
			if currentFlag.name != "" {
				currentFlag.function(currentArguments...)
				currentArguments = make([]string, 0)
			}
			for flag := range flags {
				for flagType := range flags[flag].flags {
					if arguments[argument] == flags[flag].flags[flagType] {
						if flags[flag].requiresTrailing == false {
							flags[flag].function("None")
							if flags[flag].name == "help" {
								os.Exit(0)
							}
						} else {
							currentFlag = flags[flag]
						}
					}
					break
				}
			}
		} else {
			currentArguments = append(currentArguments, arguments[argument])
		}
	}
	if currentFlag.requiresTrailing == true {
		currentFlag.function(currentArguments...)
	}
	return newGopier
}
