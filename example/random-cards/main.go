package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	card, _, _ := client.RandomCards.One()
	fmt.Printf("Name: %v\nType: %v\nRace: %v\nDescription:\n%v\n",
		card.Name, card.Type, card.Race, card.Description)
}
