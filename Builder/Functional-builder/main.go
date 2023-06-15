package main

import "fmt"

type Character struct {
  name, guild string
}

type characterModification func(*Character)

type CharacterBuilder struct {
  actions []characterModification
}

func (b *CharacterBuilder) Build() *Character {
	p := Character{}
	for _, a := range b.actions {
	  a(&p)
	}
	return &p
  }

func (b *CharacterBuilder) Name(name string) *CharacterBuilder {
  b.actions = append(b.actions, func(p *Character) {
    p.name = name
  })
  return b
}


func (b *CharacterBuilder) Guild(guild string) *CharacterBuilder {
  b.actions = append(b.actions, func(p *Character) {
    p.guild = guild
  })
  return b
}

func main() {

  cb := CharacterBuilder{}
  character := cb.Name("Diego").Guild("Shadow").Build()
  fmt.Println(*character)

}