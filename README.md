# go-yugioh

[![GoDoc](https://godoc.org/github.com/Doarakko/go-yugioh?status.svg)](https://pkg.go.dev/github.com/Doarakko/go-yugioh/yugioh)
[![Go Report Card](https://goreportcard.com/badge/github.com/Doarakko/go-yugioh)](https://goreportcard.com/report/github.com/Doarakko/go-yugioh)

go-yugioh is a Go client library for accessing the [Yu-Gi-Oh! API by YGOPRODeck v6](https://db.ygoprodeck.com/api-guide/).

## Feature

- Cards
  - Get card information with some prameters.
- Card Sets
  - Get card sets.
- Random Card
  - Get one random card.
- Card Archetypes
  - Get card archetypes.

## Install

```
$ go get -u github.com/Doarakko/go-yugioh/yugioh
```

## Documentation

Check [GoDoc](https://pkg.go.dev/github.com/Doarakko/go-yugioh/yugioh) and [API document](https://db.ygoprodeck.com/api-guide/).

## Usage

Please be carefull on the API rate limit.

> Rate Limiting on the API is enabled. The rate limit is 20 requests per 1 second. If you exceed this, you are blocked from accessing the API for 1 hour. We will monitor this rate limit for now and adjust accordingly.  
> [Yu-Gi-Oh! API Guide - YGOPRODECK](https://db.ygoprodeck.com/api-guide/)

### Cards

Get card information with some parameters.

> Note: Search target of `Q` in `CardsListOptions` is only card name.

```go
func main() {
	client := yugioh.NewClient()
	cards, _, err := client.Cards.List(
		&yugioh.CardsListOptions{
			Q:         "dragon",
			Type:      "Fusion Monster",
			Attribute: "light",
			Atk:       "gt3000",
			Misc:      "yes",
		},
	)

	for _, card := range cards[:5] {
		fmt.Printf("\"%s\" has been viewed %d times.\n", card.Name, card.Misc[0].Views)
		fmt.Printf("Type: %v\nRace: %v\nAtk: %v\n%v\n\n",
			card.Type, card.Race, card.Atk, card.Description)
	}
}
```

```
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
```

### Card Sets

Print all card sets.

> Note: This api can NOT use any parameters.

```go
func main() {
	client := yugioh.NewClient()

	cardSets, _, _ := client.CardSets.List()

	fmt.Printf("Get %v card sets, below are the first five.\n", len(cardSets))
	for _, set := range cardSets[:5] {
		fmt.Printf("%v was released on %v\n", set.Name, set.ReleaseDate)
	}
}
```

```
Get 790 card sets, below are the first five.
2-Player Starter Deck: Yuya & Declan was released on 2015-05-28
2013 Collectible Tins Wave 1 was released on 2013-08-30
2013 Collectible Tins Wave 2 was released on 2013-11-22
2014 Mega-Tin Mega Pack was released on 2014-08-28
2014 Mega-Tins was released on 2014-08-28
```

### Random Card

Get one random card.

> Note: This api can NOT use any parameters.

```go
func main() {
	client := yugioh.NewClient()

	card, _, _ := client.RandomCards.One()
	fmt.Printf("Name: %v\nType: %v\nRace: %v\n%v\n",
		card.Name, card.Type, card.Race, card.Description)
}
```

```
Name: Damage Polarizer
Type: Trap Card
Race: Counter
Activate only when an effect that inflicts damage is activated. Negate its activation and effect, and each player draws 1 card.
```

### Card Archetypes

Print all card Archetypes.

> Note: This api can NOT use any parameters.

```go
func main() {
	client := yugioh.NewClient()

	archetypes, _, _ := client.Archetypes.List()

	fmt.Printf("Get %v archetypes, below are the first five.\n", len(archetypes))
	for _, archetype := range archetypes[:5] {
		fmt.Printf("%v\n", archetype.Name)
	}
}
```

```
Get 319 archetypes, below are the first five.
@Ignister
ABC
Abyss Actor
Adamatia
Aesir
```

## License

MIT License
