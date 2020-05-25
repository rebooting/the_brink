package console

import (
	"fmt"

	"the_brink/party"
	"the_brink/world"

	"github.com/fatih/color"
	"github.com/eiannone/keyboard"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func NewMapConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 2)
	actions[0] = "Arrow | WASD Keys for Movement"
	actions[1] = "x | q to exit"

	console.Actions = actions
	return console
}


func DisplayMapConsole(gameWorld *world.World, playerParty *party.Party) {

	console := NewMapConsole()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// Main Menu Loop
	menuLoop:
	for {
		
		// Generate Map fields
		terminal := ""

		// Print from top left to bottom right
		for y := gameWorld.YMax - 1; y >= 0 ; y-- {
			line := []rune{}
			for x := 0; x < gameWorld.XMax; x++ {
				tile := world.Tile{
					X: x,
					Y: y,
				}
				line = append(line, gameWorld.Tiles[tile])
			}
			terminal += string(line) + "\n"
		}

		// Diplay
		color.Green("%s", trim)
		color.Cyan("%s", terminal)
		color.Green("%s", trim)
		console.DisplayActions()

		// char := console.ChooseKey()

		// fmt.Printf("\n\nKey pressed: %s\n\n", string(char))
		// fmt.Printf("\n\nRune pressed: %s\n\n", string(rune(char)))


		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
        if key == keyboard.KeyEsc {
			break
		}

		switch {
			// Move Left
			case key == keyboard.KeyArrowLeft, char == 'a':
				playerParty.Move(-1,0)
				gameWorld.UpdateMap()

			// Move Right
			case key == keyboard.KeyArrowRight, char == 'd':
				playerParty.Move(1,0)
				gameWorld.UpdateMap()

			// Move Up
			case key == keyboard.KeyArrowUp, char == 'w':
				playerParty.Move(0,1)
				gameWorld.UpdateMap()

			// Move Down
			case key == keyboard.KeyArrowDown, char == 's':
				playerParty.Move(0,-1)
				gameWorld.UpdateMap()

			// Exit
			case key == keyboard.KeyEsc, key == keyboard.KeyCtrlC,  key == keyboard.KeyCtrlD, char == 'q':
				break menuLoop
		}
	}
}