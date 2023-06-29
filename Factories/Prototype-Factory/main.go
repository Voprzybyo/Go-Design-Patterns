package main

import "fmt"

type Character struct {
	name   string
	guild  string
	weapon string
}

const (
	Shadow = iota
	Mecenary
)

func NewCharacter(guild int) *Character {
	switch guild {
	case Shadow:
		return &Character{"", "Shadow", "Sword"}
	case Mecenary:
		return &Character{"", "Mecenary", ""}
	default:
		panic("Not supported guild!")
	}
}

func (c *Character) WhoAreYou() {
	fmt.Printf("I am %s, I am %s and a use %s\n", c.name, c.guild, c.weapon)
}

func main() {

	diego := NewCharacter(Shadow)
	diego.name = "Diego"
	wolf := NewCharacter(Mecenary)
	wolf.name = "Wolf"
	wolf.weapon = "Bow"

	diego.WhoAreYou()
	wolf.WhoAreYou()

}