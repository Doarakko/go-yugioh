package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	card, _, _ := client.RandomCards.One()
	fmt.Printf("%v\n%v %v\n[Description]\n%v\n",
		card.Name, card.Type, card.Race, card.Description)
}
