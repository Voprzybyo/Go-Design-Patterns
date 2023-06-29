package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Place struct {
	Gate, Camp, Colony string
}

type Character struct {
	Name    string
	Place   *Place
	Friends []string
}

func (c *Character) DeepCopy() *Character {

	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(c)

	//fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Character{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	diego := Character{"Diego",
		&Place{"Main Gate", "Olc Camp", "The Colony"},
		[]string{"Lester", "Milten"}}

	milten := diego.DeepCopy()
	milten.Name = "Milten"
	milten.Place.Camp = "Old Camp"
	milten.Friends = append(milten.Friends, "Gorn")

	fmt.Println(diego, diego.Place)
	fmt.Println(milten, milten.Place)
}
