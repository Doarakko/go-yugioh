# go-yugioh

[![GoDoc](https://godoc.org/github.com/Doarakko/go-yugioh?status.svg)](https://pkg.go.dev/github.com/Doarakko/go-yugioh/yugioh)
[![Go Report Card](https://goreportcard.com/badge/github.com/Doarakko/go-yugioh)](https://goreportcard.com/report/github.com/Doarakko/go-yugioh)
![golangci-lint](https://github.com/Doarakko/go-yugioh/workflows/golangci-lint/badge.svg)

go-yugioh is a Go client library for accessing the [Yu-Gi-Oh! API by YGOPRODeck](https://db.ygoprodeck.com/api-guide/) v7.

## Feature

- Cards
  - Get card information with some prameters.
- Card Sets
  - Get card sets.
- Card Set Info
  - Get card information about set name, rarity and price from code
- Random Card
  - Get one random card.
- Card Archetypes
  - Get card archetypes.
- Check DB version

## Install

```bash
go get -u github.com/Doarakko/go-yugioh/yugioh
```

## Documentation

Check [GoDoc](https://pkg.go.dev/github.com/Doarakko/go-yugioh/yugioh) and [API document](https://db.ygoprodeck.com/api-guide/).

## Usage

Please be carefull on the API rate limit.

> Rate Limiting on the API is enabled. The rate limit is 20 requests per 1 second. If you exceed this, you are blocked from accessing the API for 1 hour. We will monitor this rate limit for now and adjust accordingly.  
> [Yu-Gi-Oh! API Guide - YGOPRODECK](https://db.ygoprodeck.com/api-guide/)

## Example

[Here](https://github.com/Doarakko/go-yugioh/tree/master/example) is the example code.

```go
client := yugioh.NewClient()
cards, _, err := client.Cards.List(
	&yugioh.CardsListOptions{
		Q:         "dragon",
		Type:      "Fusion Monster",
		Attribute: "light",
		Atk:       "gt3000",
		Misc:      "yes",
		Num:       5,
		Offset:    0,
	},
)
...

for _, card := range cards.Data {
	fmt.Printf("\"%s\" has been viewed %d times.\n", card.Name, card.Misc[0].Views)
	fmt.Printf("Type: %v\nRace: %v\nAtk: %v\n%v\n\n",
		card.Type, card.Race, card.Atk, card.Description)
}
fmt.Printf("Nex page is %s\n", cards.Meta.NextPage)
```

```text
"A-to-Z-Dragon Buster Cannon" has been viewed 50305 times.
Type: Fusion Monster
Race: Machine
Atk: 4000
"ABC-Dragon Buster" + "XYZ-Dragon Cannon"
Must be Special Summoned (from your Extra Deck) by banishing cards you control with the above original names, and cannot be Special Summoned by other ways. (You do not use "Polymerization".) During either player's turn, when your opponent activates a Spell/Trap Card, or monster effect: You can discard 1 card; negate the activation, and if you do, destroy that card. During either player's turn: You can banish this card, then target 1 each of your banished "ABC-Dragon Buster", and "XYZ-Dragon Cannon"; Special Summon them.

"Armed Dragon Catapult Cannon" has been viewed 28559 times.
Type: Fusion Monster
Race: Machine
Atk: 3500
"VWXYZ-Dragon Catapult Cannon" + "Armed Dragon LV7"
Must first be Special Summoned (from your Extra Deck) during a Duel you Special Summoned both the above cards, by banishing the above cards from your field and/or GY. (You do not use "Polymerization".) Your opponent cannot activate cards or effects with the same name as any banished card. Once per turn, during your opponent's turn (Quick Effect): You can banish 1 card from your Deck or Extra Deck, face-up; banish all cards your opponent controls and in their GY.
...
Nex page is https://db.ygoprodeck.com/api/v7/cardinfo.php?atk=gt3000&attribute=light&fname=dragon&misc=yes&num=5&type=Fusion+Monster&offset=5
```

## License

MIT License
