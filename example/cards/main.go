package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()
	cards, _, _ := client.Cards.List(
		&yugioh.CardsListOptions{
			KeyWord: "dragon", Type: "Fusion Monster", Attribute: "light"})

	for _, card := range cards {
		fmt.Printf("Name: %v\nType: %v\nRace: %v\nDescription:\n%v\n\n",
			card.Name, card.Type, card.Race, card.Description)
	}
}
