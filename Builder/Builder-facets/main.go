package main

import "fmt"

type Character struct {
	name, camp, guild string
	month, year       int
}

type CharacterBuilder struct {
	character *Character
}

type CharacterDataBuilder struct {
	CharacterBuilder
}

type CharacterAgeBuilder struct {
	CharacterBuilder
}

func NewCharacterBuilder() *CharacterBuilder {
	return &CharacterBuilder{&Character{}}
}

func (cb *CharacterBuilder) Build() *Character {
	return cb.character
}

func (cb *CharacterBuilder) Data() *CharacterDataBuilder {
	return &CharacterDataBuilder{*cb}
}

func (cb *CharacterBuilder) Age() *CharacterAgeBuilder {
	return &CharacterAgeBuilder{*cb}
}

func (cdb *CharacterDataBuilder) Name(name string) *CharacterDataBuilder {
	cdb.character.name = name
	return cdb
}

func (cdb *CharacterDataBuilder) Camp(campName string) *CharacterDataBuilder {
	cdb.character.camp = campName
	return cdb
}

func (cdb *CharacterDataBuilder) Guild(guildName string) *CharacterDataBuilder {
	cdb.character.guild = guildName
	return cdb
}

func (cab *CharacterAgeBuilder) Year(year int) *CharacterAgeBuilder {
	cab.character.year = year
	return cab
}

func (cab *CharacterAgeBuilder) Month(month int) *CharacterAgeBuilder {
	cab.character.month = month
	return cab
}

func main() {
	ch := NewCharacterBuilder()
	ch.Data().Name("Diego").Camp("Old Camp").Guild("Shadow").Age().Year(2001).Month(3)

	diego := ch.Build()
	fmt.Println(diego)
}
