package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	cardSets, _, err := client.CardSets.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Get %v card sets, below are the first five.\n", len(cardSets))
	for _, set := range cardSets[:5] {
		fmt.Printf("%v was released on %v\n", set.Name, set.ReleasedDate)
	}
}
