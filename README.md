# go-yugioh

[![GoDoc](https://godoc.org/github.com/Doarakko/go-yugioh?status.svg)](https://godoc.org/github.com/Doarakko/go-yugioh)
[![Go Report Card](https://goreportcard.com/badge/github.com/Doarakko/go-yugioh)](https://goreportcard.com/report/github.com/Doarakko/go-yugioh)

go-yugioh is a Go client library for accessing the [Yu-Gi-Oh! API by YGOPRODeck](https://db.ygoprodeck.com/api-guide/) v5.

## Feature

- Cards
  - Get card information with some prameters.
- Card Sets
  - Get card sets.
- Random Card
  - Get one random card.

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

Get card information with some parameters.

> Note: Search target of `KeyWord` in `CardsListOptions` is only card name.

```
func main() {
	client := yugioh.NewClient()
	cards, _, _ := client.Cards.List(
		&yugioh.CardsListOptions{
			KeyWord: "dragon", Type: "Fusion Monster", Attribute: "light"})

	for _, card := range cards {
		fmt.Printf("Name: %v\nType: %v\nRace: %v\nDescription:\n%v\n\n",
			card.Name, card.Type, card.Race, card.Description)
	}
}
```

```
Name: A-to-Z-Dragon Buster Cannon
Type: Fusion Monster
Race: Machine
Description:
"ABC-Dragon Buster" + "XYZ-Dragon Cannon"
Must be Special Summoned (from your Extra Deck) by banishing cards you control with the above original names, and cannot be Special Summoned by other ways. (You do not use "Polymerization".) During either player's turn, when your opponent activates a Spell/Trap Card, or monster effect: You can discard 1 card; negate the activation, and if you do, destroy that card. During either player's turn: You can banish this card, then target 1 each of your banished "ABC-Dragon Buster", and "XYZ-Dragon Cannon"; Special Summon them.

Name: ABC-Dragon Buster
Type: Fusion Monster
Race: Machine
Description:
"A-Assault Core" + "B-Buster Drake" + "C-Crush Wyvern"
Must first be Special Summoned (from your Extra Deck) by banishing the above cards you control and/or from your Graveyard. (You do not use "Polymerization".) Once per turn, during either player's turn: You can discard 1 card, then target 1 card on the field; banish it. During your opponent's turn: You can Tribute this card, then target 3 of your banished LIGHT Machine-Type Union monsters with different names; Special Summon them (this is a Quick Effect).
...
```

### Card Sets

Print all card sets.

> Note: This api can NOT use any parameters.

```
func main() {
	client := yugioh.NewClient()

	cardSets, _, _ := client.CardSets.List()

	for _, set := range cardSets {
		fmt.Printf("%v\n", set.Name)
	}

}
```

```
2-Player Starter Deck: Yuya & Declan
2013 Collectible Tins Wave 1
2013 Collectible Tins Wave 2
2014 Mega-Tin Mega Pack
2014 Mega-Tins
2015 Mega-Tin Mega Pack
2015 Mega-Tins
2016 Mega-Tin Mega Pack
2016 Mega-Tins
2017 Mega-Tin Mega Pack
2017 Mega-Tins
2018 Mega-Tin Mega Pack
2018 Mega-Tins
...
```

### Random Card

Get one random card.

> Note: This api can NOT use any parameters.

```
func main() {
	client := yugioh.NewClient()

	card, _, _ := client.RandomCards.One()
	fmt.Printf("Name: %v\nType: %v\nRace: %v\nDescription:\n%v\n",
		card.Name, card.Type, card.Race, card.Description)
}
```

```
Name: Formula Synchron
Type: Synchro Tuner Monster
Race: Machine
Description:
1 Tuner + 1 non-Tuner monster
When this card is Synchro Summoned: You can draw 1 card. Once per Chain, during your opponent's Main Phase, you can: Immediately after this effect resolves, Synchro Summon using this card you control (this is a Quick Effect).
```

### Card Archetypes

Print all card Archetypes.

> Note: This api can NOT use any parameters.

```
func main() {
	client := yugioh.NewClient()

	archetypes, _, _ := client.Archetypes.List()

	for _, archetype := range archetypes {
		fmt.Printf("%v\n", archetype.Name)
	}
}
```

```
@Ignister
ABC
Abyss Actor
Aesir
Aether
Alien
Alligator
Allure Queen
Ally of Justice
Altergeist
Amazoness
...
```

## License

MIT License
