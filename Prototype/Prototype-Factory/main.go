package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Place struct {
	Cabin  int
	Colony string
	Camp   string
}

type Character struct {
	Name        string
	PlaceToLive Place
}

func (c *Character) DeepCopy() *Character {

	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(c)

	d := gob.NewDecoder(&b)
	result := Character{}
	_ = d.Decode(&result)
	return &result
}

var LeftSideOfCamp = Character{"", Place{0, "The Colony", "Old Camp"}}
var RightSideOfCamp = Character{"", Place{0, "The Colony", "Old Camp"}}

func newCharacter(proto *Character, name string, cabin int) *Character {
	result := proto.DeepCopy()
	result.Name = name
	result.PlaceToLive.Cabin = cabin
	return result
}

func NewLeftSideCampCharacter(name string, cabin int) *Character {
	return newCharacter(&LeftSideOfCamp, name, cabin)
}

func NewRightSideCampCharacter(name string, cabin int) *Character {
	return newCharacter(&RightSideOfCamp, name, cabin)
}

func main() {
	diego := NewLeftSideCampCharacter("Diego", 2)
	fisk := NewRightSideCampCharacter("Fisk", 7)

	fmt.Println(diego)
	fmt.Println(fisk)
}
