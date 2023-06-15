package main

import (
	"fmt"
	"strings"
)

type message struct {
	from, to, body string
}

type MessageBuilder struct {
	message message
}

func (b *MessageBuilder) From(from string) *MessageBuilder {
	b.message.from = from
	return b
}

func (b *MessageBuilder) To(to string) *MessageBuilder {
	b.message.to = to
	return b
}

func (b *MessageBuilder) Body(body string) *MessageBuilder {
	if !strings.Contains(body, "Bread") {
		panic("No bread!")
	}
	b.message.body = body
	return b
}

func GiveLetter(message *message) {
	fmt.Println(message)
}

type build func(*MessageBuilder)

func SendMessage(action build) {
	builder := MessageBuilder{}
	action(&builder)
	GiveLetter(&builder.message)
}

func main() {
	SendMessage(func(b *MessageBuilder) {
		b.From("Diego").To("Jan").Body("Bread")
	})
}