package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Camp string
}

// func NewPerson(name string, age int) Person {
//   return Person{name, age}
// }

func NewPerson(name string, age int) *Person {
	if name != "Gomez" {
		panic("And?!")
	}
	return &Person{name, age, "Old Camp"}
}

func main() {
	// Directly...
	//p := Person{"Diego", 22}
	//fmt.Println(p)

	// Usage of constructor
	gomez := NewPerson("Gomez", 35)
	gomez.Age = 18
	fmt.Println(gomez)
}
