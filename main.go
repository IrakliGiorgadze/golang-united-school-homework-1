package main

import (
	"github.com/kyokomi/emoji"
)

func main() {
	GetMessage()
}

func GetMessage() string {
	emoji := emoji.Sprint("Hello :world_map:!")
	return emoji
}
