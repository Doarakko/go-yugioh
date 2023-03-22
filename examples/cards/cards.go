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
			Num:       5,
			Offset:    0,
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, card := range cards.Data {
		fmt.Printf("\"%s\" has been viewed %d times.\n", card.Name, card.Misc[0].Views)
		fmt.Printf("Type: %v\nRace: %v\nAtk: %v\n%v\n\n",
			card.Type, card.Race, card.Atk, card.Description)
	}
	fmt.Printf("Nex page is %s\n", cards.Meta.NextPage)
}
