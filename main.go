package main

import (
	"flag"
	"github.com/sgerhardt/HeartsOfIronModTools/NavyConv"
	"github.com/sgerhardt/HeartsOfIronModTools/Parse"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inputDirPointer := flag.String("in", "", "Directory containing unit files to parse")
	outputDirPointer := flag.String("out", "", "Directory containing output files")
	flag.Parse()

	files, err := ioutil.ReadDir(*inputDirPointer)
	if err != nil {
		panic("err reading directory:" + err.Error())
	}
	for _, file := range files {
		if !file.IsDir() {
			fileData, err := ioutil.ReadFile(*inputDirPointer + string(os.PathSeparator) + file.Name())
			if !strings.Contains(string(fileData), "navy") {
				continue
			}
			originalData := fileData
			if err != nil {
				panic("err reading file:" + err.Error())
			}
			navies := []*NavyConv.Navy{}
			naviesText := []string{}
			for {
				nextLoc := 0
				navy, nextLoc := Parse.For("navy", string(fileData[nextLoc:]))
				if nextLoc == -1 {
					break
				}
				fileData = fileData[nextLoc:]
				naviesText = append(naviesText, navy)
				navies = append(navies, NavyConv.Parse(navy))
			}

			fleets := NavyConv.ToFleets(navies)

			fleetsData := ""
			for _, fleet := range fleets {
				fleetsData += fleet.String()
			}

			// write new files with updated fleets
			err = ioutil.WriteFile(*outputDirPointer+file.Name(),
				[]byte(NavyConv.InsertFleetsIntoUnits(NavyConv.ClearOldNavies(string(originalData), naviesText), fleetsData)),
				os.ModeAppend)
			if err != nil {
				panic(err)
			}

		}
	}
}
