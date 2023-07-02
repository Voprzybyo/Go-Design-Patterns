package main

import "fmt"

type Guild string

const (
	Shadow    Guild = "Shadow"
	Mercenary       = "Mercenary"
)

type Character interface {
	WhoAreYou()
}

type character struct { // <-- lowercase
	name  string
	guild Guild
}

type mainCharacter struct {
	guild Guild
}

func (c *character) WhoAreYou() {
	fmt.Printf("I am %s, I am %s \n", c.name, c.guild)
}

func (c *mainCharacter) WhoAreYou() {
	fmt.Printf("I am ..., I am %s \n", c.guild)
}

func NewCharacter(name string, guild Guild) Character { // <-- Returning interface
	if name == "" {
		return &mainCharacter{guild}
	}
	return &character{name, guild}
}

func main() {
	diego := character{"Diego", Shadow}
	diego.WhoAreYou()
	diego.guild = Mercenary // Access to structure fields (not interface)
	diego.WhoAreYou()

	wolf := NewCharacter("Wolf", Mercenary)
	wolf.WhoAreYou()
	//wolf.guild = Shadow // Cannot do this becouse this is interface

	mainCharacter := NewCharacter("", Shadow)
	mainCharacter.WhoAreYou()
}
