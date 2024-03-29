package Parse

import "testing"
import "github.com/stretchr/testify/assert"

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
	testText := `	division= { 
			name = "5a División de Infantería"
			location = 12349 # Trelew
			division_template="Infantry Division"
			start_experience_factor=0.05
			} ` + navyText + `
air_wings = { 
	}`
	parsedText, index := For("navy", testText)
	assert.Equal(t, navyText, parsedText)
	assert.Equal(t, 2439, index)
}
