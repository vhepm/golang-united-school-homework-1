package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	var compiled = emoji.Sprint("Hello :world_map:!")
	return compiled
}
