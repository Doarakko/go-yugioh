package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	archetypes, _, _ := client.Archetypes.List()

	for _, archetype := range archetypes {
		fmt.Printf("%v\n", archetype.Name)
	}
}
