package main

import "fmt"

type Character struct {
	name   string
	guild  string
	weapon string
}

func (c *Character) WhoAreYou() {
	fmt.Printf("I am %s, I am %s and my weapon is %s\n", c.name, c.guild, c.weapon)
}
func (c *CharacterFactory) WhoAreYou() {
	//fmt.Printf("I am %s, I am %s and my weapon is %s\n", c.name, c.guild, c.weapon)
}
// Functional approach
// vvvvvvvvvvvvvvvvvvv
func NewCharacterFactory(guild string, weapon string) func(name string) *Character {
	return func(name string) *Character {
		return &Character{name, guild, weapon}
	}
}

// Structural approach
// vvvvvvvvvvvvvvvvvvv
type CharacterFactory struct {
	//name   string
	guild  string
	weapon string
}

func (f CharacterFactory) Create(name string) *Character {
	return &Character{name, f.guild, f.weapon}
} 

func NewCharacterFactory2 (guild string, weapon string) *CharacterFactory {
	return &CharacterFactory{guild, weapon}
}

func main() {
	ShadowFactory := NewCharacterFactory("Shadow", "Sword")
	MercenaryFactory := NewCharacterFactory("Mecenary", "Bow")

	Diego := ShadowFactory("Diego")
	Wolf := MercenaryFactory("Wolf")

	Diego.WhoAreYou()
	Wolf.WhoAreYou()

	mageFactory := NewCharacterFactory2("Mage", "Rune")
	Corristo := mageFactory.Create("Corristo")
	fmt.Println(Corristo)

}
