package solution

import (
	"github.com/kyokomi/emoji"
)

func main() {
	GetMessage()
}

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
