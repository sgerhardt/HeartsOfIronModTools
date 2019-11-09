package NavyConv

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	navyText := `navy = {
		name = "Armada Argentina"
		location=12364 # Buenos Aires
		base=12364 # Buenos Aires
ship = { name = "ARA La Plata" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA Los Andes" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA Almirante Brown" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG version_name = "Almirante Brown Class" } } }
ship = { name = "ARA Independencia" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG version_name = "Libertad Class" } } }
ship = { name = "ARA Libertad" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG version_name = "Libertad Class" } } }
ship = { name = "ARA Patagonia" definition = light_cruiser equipment = { light_cruiser_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA Veinticinco de Mayo" definition = light_cruiser equipment = { light_cruiser_1890 = { amount = 1 owner = ARG version_name = "25 De Mayo Class" } } }
ship = { name = "ARA Nueve de Julio" definition = light_cruiser equipment = { light_cruiser_1890 = { amount = 1 owner = ARG version_name = "9 De Julio Class" } } }
ship = { name = "ARA Buenos Aires" definition = light_cruiser equipment = { light_cruiser_1900 = { amount = 1 owner = ARG } } }
ship = { name = "ARA General Garibaldi" definition = heavy_cruiser equipment = { heavy_cruiser_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA General Belgrano" definition = heavy_cruiser equipment = { heavy_cruiser_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA General Pueyrredón" definition = heavy_cruiser equipment = { heavy_cruiser_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA General San Martín" definition = heavy_cruiser equipment = { heavy_cruiser_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA Corrientes" definition = destroyer equipment = { destroyer_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA Entre Rios" definition = destroyer equipment = { destroyer_1890 = { amount = 1 owner = ARG } } }
ship = { name = "ARA Misiones" definition = destroyer equipment = { destroyer_1890 = { amount = 1 owner = ARG } } }
		}`
	navy := Parse(navyText)
	assert.Equal(t, `"Armada Argentina"`, navy.name)
	assert.Len(t, navy.ships, 16)
	assert.Equal(t, "12364 # Buenos Aires", navy.base)
}

func TestToFleet(t *testing.T) {
	b1 := "base1#LocRef"
	b2 := "base2#LocRef2"
	smallNavy := Navy{
		base: b1,
		name: "smallNavy",
		ships: []string{
			`ship = { name = "ARA Misiones" definition = destroyer equipment = { destroyer_1890 = { amount = 1 owner = ARG } } }`,
			`ship = { name = "ARA Entre Rios" definition = destroyer equipment = { destroyer_1890 = { amount = 1 owner = ARG } } }`,
		},
	}
	largeNavy := Navy{
		base: b1,
		name: "smallNavy",
		ships: []string{
			`ship = { name = "ARA General San Martín" definition = heavy_cruiser equipment = { heavy_cruiser_1890 = { amount = 1 owner = ARG } } }`,
			`ship = { name = "ARA General Belgrano" definition = heavy_cruiser equipment = { heavy_cruiser_1890 = { amount = 1 owner = ARG } } }`,
			`ship = { name = "ARA Los Andes" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG } } }`,
			`ship = { name = "ARA Almirante Brown" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG version_name = "Almirante Brown Class" } } }`,
		},
	}
	diffBaseNavy := Navy{
		base: b2,
		name: "tinyNavy",
		ships: []string{
			`ship = { name = "ARA Libertad" definition = battleship equipment = { battleship_1890 = { amount = 1 owner = ARG version_name = "Libertad Class" } } }`,
		},
	}
	fleets := ToFleets([]*Navy{&smallNavy, &largeNavy, &diffBaseNavy})
	assert.Len(t, fleets, 2)
	for _, fleet := range fleets {
		require.NotEmpty(t, fleet.name)
		if fleet.base == b1 {
			require.Len(t, fleet.taskForce, 2)
			assert.Equal(t, fleet.taskForce[0].base, b1)
			assert.Len(t, fleet.taskForce[0].ships, 2)
			assert.Equal(t, fleet.taskForce[1].base, b1)
			assert.Len(t, fleet.taskForce[1].ships, 4)
		} else if fleet.base == b2 {
			assert.Len(t, fleet.taskForce, 1)
			assert.Equal(t, fleet.taskForce[0].base, b2)
			assert.Len(t, fleet.taskForce[0].ships, 1)
		}
	}
}

func TestClearOldNavies(t *testing.T) {
	navy1 := `    navy = {
        name = "1. Schlachtschwadron"
        location = 241#Wilhelmshaven
        base = 241#Wilhelmshaven
ship = { name = "SMS Nassau" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER } } }
ship = { name = "SMS Westfalen" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER } } }
ship = { name = "SMS Rheinland" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER } } }
ship = { name = "SMS Posen" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER } } }
ship = { name = "SMS Helgoland" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER version_name = "Helgoland Class" } } }
ship = { name = "SMS Ostfriesland" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER version_name = "Helgoland Class" } } }
ship = { name = "SMS Thüringen" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER version_name = "Helgoland Class" } } }
ship = { name = "SMS Oldenburg" definition = battleship equipment = { battleship_1906 = { amount = 1 owner = GER version_name = "Helgoland Class" } } }

    }

`

	navy2 := `    navy = {
        name = "SMS Dresden"
        location = 241#Wilhelmshaven
        base = 241#Wilhelmshaven
ship = { name = "SMS Dresden" definition = light_cruiser equipment = { light_cruiser_1906 = { amount = 1 owner = GER version_name = "Dresden Class" } } }

    }`

	placeHolderText := "some text\n"
	placeholderText2 := "some other text \n"
	input := placeHolderText + navy1 + navy2 + placeholderText2

	naviesRemoved := ClearOldNavies(input, []string{navy1, navy2})

	assert.Equal(t, placeHolderText+placeholderText2, naviesRemoved)
}
