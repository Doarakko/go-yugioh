package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	card, _, err := client.RandomCards.One()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %v\nType: %v\nRace: %v\n%v\n",
		card.Name, card.Type, card.Race, card.Description)
}
