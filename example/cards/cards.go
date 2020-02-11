package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()
	cards, _, err := client.Cards.List(
		&yugioh.CardsListOptions{
			Q:         "dragon",
			Type:      "Fusion Monster",
			Attribute: "light",
			Atk:       "gt3000",
			Misc:      "yes",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, card := range cards[:5] {
		fmt.Printf("\"%s\" has been viewed %d times.\n", card.Name, card.Misc[0].Views)
		fmt.Printf("Type: %v\nRace: %v\nAtk: %v\n%v\n\n",
			card.Type, card.Race, card.Atk, card.Description)
	}
}
