package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//Flag is creating a simple struct for holding information about the flags that can be used to configure gopyCode.
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

/*
	Setting the default flags that can be used in the configuration of gopyCode.
	Re also used as the template for the help method of the flags.
	On -h all flags below are parsed and display as a quick overview menu.
*/
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
		description: "Define the extennsions, gopyCode should look for like \n		-> -ex .java .go\n	 Or just take all with\n		 -> -ex .",
		function: setExtensions,
	})

	flags = append(flags, Flag{
		name:             "timed",
		flags:            []string{"-t", "--timer", "-timed", "--timed"},
		requiresTrailing: false,
		description: "Let gopyCode tell you, how long it had to work like \n		-> -t",
		function: setTimer,
	})
}

func setTimer(none ...string) {
	newGopier.timed = true
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
	var allFlags []string
	for i := range flags {
		if flags[i].name != "help" {
			allFlags = append(allFlags, flags[i].flags[0])
			manual += fmt.Sprintf("*  %v	%v\n	 %v\n", flags[i].flags, flags[i].name, flags[i].description)
		}
	}
	fmt.Println(fmt.Sprintf("\ngopyCode [%v]", strings.Join(allFlags, ", ")))
	fmt.Println("\n" + manual)
}

func setFlags() {
	addDefaultFlags()
	addFlags()
}

var newGopier Gopier

func checkForHelp(arguemnts []string) {
	helpFlags := []string{"-h", "--h", "-help"}
	for arg := range arguemnts {
		for x := range helpFlags {
			if arguemnts[arg] == helpFlags[x] {
				getGopyCodeManual("")
				os.Exit(0)
			}
		}
	}
}

func parseFlags(arguments []string) Gopier {
	checkForHelp(arguments)
	newGopier = Gopier{}
	var currentFlag Flag
	var currentArguments []string
	for argument := range arguments {
		found := false
		if strings.HasPrefix(arguments[argument], "-") {
			if currentFlag.name != "" {
				if len(currentArguments) != 0 {
					currentFlag.function(currentArguments...)
					currentArguments = make([]string, 0)
				} else {
					panic(fmt.Sprintf("%v [%v] is in need of arguments", currentFlag.name, strings.Join(currentFlag.flags, ", ")))
				}
			}
			for flag := range flags {
				allFlags := flags[flag].flags
				for flagType := range allFlags {
					if arguments[argument] == allFlags[flagType] {
						found = true
						if flags[flag].requiresTrailing == false {
							flags[flag].function("None")
							if flags[flag].name == "help" {
								os.Exit(0)
							}
							currentFlag = Flag{}
						} else {
							currentFlag = flags[flag]
						}
					}
				}

			}
			if found == false {
				fmt.Println(fmt.Sprintf("{\n	%v doesn't seem to be a known flag\n 	Use -h or --help for help with gopyCode\n}", arguments[argument]))
				os.Exit(0)
			}
		} else {
			currentArguments = append(currentArguments, arguments[argument])
		}
	}
	if currentFlag.name != "" {
		currentFlag.function(currentArguments...)
	}
	return newGopier
}
