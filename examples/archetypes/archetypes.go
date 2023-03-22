package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	archetypes, _, err := client.Archetypes.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Get %v archetypes, below are the first five.\n", len(archetypes))
	for _, archetype := range archetypes[:5] {
		fmt.Printf("%v\n", archetype.Name)
	}
}
