package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	c "github.com/sergivillar/rock-paper-scissors/config"
	p "github.com/sergivillar/rock-paper-scissors/player"
)

type game struct {
	players []p.Player
}

func main() {
	p1, err := p.Create("Player")
	if err != nil {
		log.Println(err)
	}
	p2, err := p.Create("CPU")
	if err != nil {
		log.Println(err)
	}

	g := game{}
	g.addPlayer(p1)
	g.addPlayer(p2)

	fmt.Printf(`Game starts. Please select one:
1 - %v
2 - %v
3 - %v
0 - Exit program
`, strings.Title(c.GameOptions[0]),
		strings.Title(c.GameOptions[1]),
		strings.Title(c.GameOptions[2]))

	for {
		var option int
		fmt.Print("Enter option: ")
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Error while reading input.")
			os.Exit(1)
		}
		if option == 0 {
			fmt.Println("See you next time.")
			os.Exit(1)
		}

		if option > len(c.GameOptions) {
			fmt.Println("Invalid option. Choose another one")
			continue
		}

		fmt.Println(g.play(c.GameOptions[option-1]))
	}
}

func (g *game) addPlayer(p p.Player) {
	g.players = append(g.players, p)
}

func (g game) play(playerOption string) string {
	cpuOption := p.RockPaperScissor()

	var result string
	if playerOption == cpuOption {
		result = "Draw"
	} else if c.Rules[playerOption] == cpuOption {
		result = fmt.Sprintf("%v won", g.players[0].Name)
	} else {
		result = fmt.Sprintf("%v won", g.players[1].Name)
	}
	return result
}
