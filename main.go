package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	var compiled = emoji.Sprint("Hello :world_map:!")
	fmt.Println(compiled)
}
