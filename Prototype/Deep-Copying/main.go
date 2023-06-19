package main

import "fmt"

type Place struct {
	Gate, Camp, Colony string
}

type Character struct {
	Name    string
	Place *Place
}

func main() {
	diego := Character{"Diego", &Place{"Main Gate", "Old Camp", "The Colony"}}

	wolf := diego

	// Shallow copy
	wolf.Name = "Wolf" // ok
	//wolf.Place.Gate = "321 Baker St" // Not ok (pointer)

	fmt.Println(diego.Name, diego.Place)
	fmt.Println(wolf.Name, wolf. Place)

	wolf2 := diego
	wolf2.Place = &Place{
		diego.Place.Gate,
		diego.Place.Camp,
		diego.Place.Colony}

	wolf2.Name = "Wolf"          // Ok
	wolf2.Place.Camp = "New Camp"// Ok
	wolf2.Place.Gate = "Center"  // Ok

	fmt.Println(diego.Name, diego.Place)
	fmt.Println(wolf2.Name, wolf2.Place)
}
