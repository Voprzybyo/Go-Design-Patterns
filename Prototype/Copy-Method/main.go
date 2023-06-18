package main

import "fmt"

type Place struct {
	Gate, Camp, Colony string
}

type Character struct {
	Name    string
	Place   *Place
	Friends []string
}

func (a *Place) DeepCopy() *Place {
	return &Place{a.Gate, a.Camp, a.Colony}
}

func (c *Character) DeepCopy() *Character {
	cp := *c
	cp.Place = cp.Place.DeepCopy()
	copy(cp.Friends, c.Friends)
	return &cp
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

	// index := -1
	// for i, v := range milten.Friends {
	// 	if v == "Milten" {
	// 		index = i
	// 	}
	// }

	// milten.Friends = append(milten.Friends[:index], milten.Friends[index+1:]...)
	// fmt.Println(milten.Friends)
}
