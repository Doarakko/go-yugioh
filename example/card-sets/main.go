package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	cardSets, _, _ := client.CardSets.List()

	for _, set := range cardSets {
		fmt.Printf("%v\n", set.Name)
	}

}
