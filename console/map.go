package console

import (
	"fmt"
	"os"
	"os/exec"

	"the_brink/party"
	"the_brink/world"

	"github.com/fatih/color"
)


var trim string = "^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^-^\n"

func NewMapConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 5)
	actions[0] = "Arrow | WASD Keys for Movement"
	actions[1] = "x | q to exit"

	console.Actions = actions
	return console
}


func DisplayMapConsole(gameWorld *world.World, playerParty *party.Party) {

	console := NewMapConsole()

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

		_ = console.ChooseAction()
	
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		fmt.Printf("\n\nKey pressed: %s\n\n", string(char))
		fmt.Printf("\n\nRune pressed: %s\n\n", string(rune(char)))

		if err != nil {
			fmt.Pritln(err.Error())
		}


		switch rune(char) {
			// Left
			case "Left":
				playerParty.Move(-1,0)
				gameWorld.UpdateMap()

			// Right
			case "Right":
				playerParty.Move(1,0)
				gameWorld.UpdateMap()

			// Up
			case "Up":
				playerParty.Move(0,1)
				gameWorld.UpdateMap()

			// Down
			case "Down":
				playerParty.Move(0,-1)
				gameWorld.UpdateMap()

			// Exit
			case "Exit":
				break menuLoop
		}

		i++
	}
}