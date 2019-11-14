# go-yugioh

[![GoDoc](https://godoc.org/github.com/Doarakko/go-yugioh?status.svg)](https://godoc.org/github.com/Doarakko/go-yugioh)
[![Go Report Card](https://goreportcard.com/badge/github.com/Doarakko/go-yugioh)](https://goreportcard.com/report/github.com/Doarakko/go-yugioh)

go-yugioh is a Go client library for accessing the [Yu-Gi-Oh! API by YGOPRODeck](https://db.ygoprodeck.com/api-guide/) v5.

## Feature

- Cards
- Card Sets
- Random Card

## Install

```
$ go get github.com/Doarakko/go-yugioh
```

## Documentation

Check [GoDoc](https://godoc.org/github.com/Doarakko/go-yugioh) and [API document](https://db.ygoprodeck.com/api-guide/).

## Usage

Please be carefull on the API rate limit.

> Rate Limiting on the API is enabled. The rate limit is 20 requests per 1 second. If you exceed this, you are blocked from accessing the API for 1 hour. We will monitor this rate limit for now and adjust accordingly.  
> [Yu-Gi-Oh! API Guide - YGOPRODECK](https://db.ygoprodeck.com/api-guide/)

### Cards

> Note: Search target of `KeyWord` in `CardsListOptions` is only card name.

```
import (
	"fmt"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()
	cards, _, _ := client.Cards.List(&yugioh.CardsListOptions{KeyWord: "dragon"})

	for _, card := range cards {
		fmt.Printf("%v\n%v %v\n[Description]\n%v\n\n",
			card.Name, card.Type, card.Race, card.Description)
	}
}
```

```
A Wingbeat of Giant Dragon
Spell Card Normal
[Description]
Return 1 Level 5 or higher Dragon-Type monster you control to the hand, and if you do, destroy all Spell and Trap Cards on the field.

A-to-Z-Dragon Buster Cannon
Fusion Monster Machine
[Description]
"ABC-Dragon Buster" + "XYZ-Dragon Cannon"
Must be Special Summoned (from your Extra Deck) by banishing cards you control with the above original names, and cannot be Special Summoned by other ways. (You do not use "Polymerization".) During either player's turn, when your opponent activates a Spell/Trap Card, or monster effect: You can discard 1 card; negate the activation, and if you do, destroy that card. During either player's turn: You can banish this card, then target 1 each of your banished "ABC-Dragon Buster", and "XYZ-Dragon Cannon"; Special Summon them.

ABC-Dragon Buster
Fusion Monster Machine
[Description]
"A-Assault Core" + "B-Buster Drake" + "C-Crush Wyvern"
Must first be Special Summoned (from your Extra Deck) by banishing the above cards you control and/or from your Graveyard. (You do not use "Polymerization".) Once per turn, during either player's turn: You can discard 1 card, then target 1 card on the field; banish it. During your opponent's turn: You can Tribute this card, then target 3 of your banished LIGHT Machine-Type Union monsters with different names; Special Summon them (this is a Quick Effect).

...
```

### Card Sets

### Random Card

## License

MIT License
