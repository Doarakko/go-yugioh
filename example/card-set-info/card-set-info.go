package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	cardSetInfo, _, err := client.CardSetInfo.One(
		&yugioh.CardSetInfoOneOptions{Code: "SDY-046"},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("set name: %s\ncard rarity: %s\n",
		cardSetInfo.Name,
		cardSetInfo.Rarity,
	)
}
