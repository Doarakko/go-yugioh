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
			Atk:       "gte1200",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, card := range cards[:5] {
		fmt.Printf("Name: %v\nType: %v\nRace: %v\nAtk: %v\n%v\n\n",
			card.Name, card.Type, card.Race, card.Atk, card.Description)
	}
}
