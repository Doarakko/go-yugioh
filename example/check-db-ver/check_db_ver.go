package main

import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	dbInformation, _, err := client.CheckDBVersion.One()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("database version: %s last updated: %s\n", dbInformation.Version, dbInformation.LastUpdated)
}
