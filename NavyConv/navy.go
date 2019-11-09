package NavyConv

import (
	"fmt"
	"strconv"
	"strings"

	genericParse "github.com/sgerhardt/HeartsOfIronModTools/Parse"
)

type Navy struct {
	base  string
	name  string
	ships []string
}

type fleet struct {
	name      string
	base      string
	taskForce []Navy
}

func Parse(navy string) *Navy {
	n := strings.Index(navy, `name`)
	nameStart := strings.Index(navy[n:], `"`) + n
	name := navy[nameStart : strings.Index(navy[nameStart+1:], `"`)+nameStart+2]

	b := strings.Index(navy, `base`)
	baseStart := strings.Index(navy[b:], `=`) + b + 1
	baseName := ""
	for _, char := range strings.TrimSpace(navy[baseStart:]) {
		i, err := strconv.Atoi(string(char))
		if err != nil {
			// If there's a comment, add it until a new-line
			endLine := strings.Index(navy[baseStart:], "\n")
			commentIndex := strings.IndexAny(navy[baseStart:baseStart+endLine], "#")
			if commentIndex != -1 {
				fmt.Printf(name)
				baseName += " " + navy[baseStart+commentIndex:baseStart+endLine]
			}
			break
		}
		baseName += strconv.Itoa(i)
	}

	ships := []string{}
	for {
		nextLoc := 0
		ship, nextLoc := genericParse.For("ship", navy[nextLoc:])
		if nextLoc == -1 {
			break
		}
		navy = navy[nextLoc:]
		ships = append(ships, ship)
		fmt.Printf("ship Parsed:%+v\n", ship)
	}

	return &Navy{base: baseName, name: name, ships: ships}
}

func ToFleets(navies []*Navy) []fleet {
	// combine all navies in the same base into one fleet
	baseToNavies := map[string][]Navy{}
	for _, navy := range navies {
		baseToNavies[navy.base] = append(baseToNavies[navy.base], *navy)
	}

	fleets := []fleet{}
	for base, navies := range baseToNavies {

		// Fleet name will default to the largest navy in a base
		largestNavy := 0
		largestNavyName := ""
		for _, navy := range navies {
			if len(navy.ships) > largestNavy {
				largestNavyName = navy.name
			}
		}

		fleets = append(fleets, fleet{
			name:      getFleetName(largestNavyName),
			base:      base,
			taskForce: navies,
		})

	}
	fmt.Printf("Fleets are: %+v\n", fleets)
	return fleets
}

func getFleetName(navyName string) string {
	lastDblQuote := strings.LastIndex(navyName, `"`)
	if lastDblQuote != -1 {
		navyName := navyName[:lastDblQuote]
		navyName += ` Fleet"`
	}
	navyName += ` Fleet"`
	return strings.TrimSpace(navyName)
}

func (f *fleet) String() string {
	fleetStr := "\nfleet = {\n"

	fleetStr += "	name = " + f.name + "\n"
	fleetStr += "	naval_base = " + f.base + "\n"
	for _, f := range f.taskForce {
		fleetStr += "	task_force = {\n"
		fleetStr += "		name = " + f.name + "\n"
		fleetStr += "		location = " + f.base + "\n"
		for _, ship := range f.ships {
			fleetStr += "		" + ship + "\n"
		}
		fleetStr += "	}\n"
	}
	fleetStr += "}\n"
	return fleetStr
}

func InsertFleetsIntoUnits(input string, fleets string) string {
	return strings.ReplaceAll(input, `units = {`, "units = {\n"+fleets)
}

func ClearOldNavies(input string, naviesText []string) string {
	for _, navy := range naviesText {
		input = strings.ReplaceAll(input, navy, "")
	}
	return input
}
