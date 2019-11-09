package main

import (
	"flag"
	"github.com/sgerhardt/HeartsOfIronModTools/NavyConv"
	"github.com/sgerhardt/HeartsOfIronModTools/Parse"
	"io/ioutil"
	"os"
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
			workingData, err := ioutil.ReadFile(*inputDirPointer + string(os.PathSeparator) + file.Name())
			originalData := workingData
			if err != nil {
				panic("err reading file:" + err.Error())
			}

			//naviesText := map[string]struct{}{}
			navies := []*NavyConv.Navy{}
			naviesText := []string{}
			for {
				nextLoc := 0
				navy, nextLoc := Parse.For("navy", string(workingData[nextLoc:]))
				if nextLoc == -1 {
					break
				}
				workingData = workingData[nextLoc:]
				naviesText = append(naviesText, navy)
				navies = append(navies, NavyConv.Parse(navy))
			}

			fleets := NavyConv.ToFleets(navies)

			fleetsData := ""
			for _, fleet := range fleets {
				fleetsData += fleet.String()
			}
			updatedData := NavyConv.ClearOldNavies(string(originalData), naviesText)

			updatedData = NavyConv.InsertFleetsIntoUnits(string(originalData), fleetsData)

			// write new files with updated fleets
			err = ioutil.WriteFile(*outputDirPointer+file.Name(), []byte(updatedData), os.ModeAppend)
			if err != nil {
				panic(err)
			}

		}
	}
}
