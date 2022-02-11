package main

import (
	"github.com/kyokomi/emoji"
)

func main() {
	GetMessage()
}

func GetMessage() string {

	emoji := emoji.Sprintf("Hello :world_map:!")
	return emoji
}
