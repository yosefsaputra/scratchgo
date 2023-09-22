package main

import (
	"fmt"

	"github.com/yosefsaputra/scratchgo/debug"
)

func main() {
	fmt.Println("Hello World!")

	// debug usage
	mystr := "test"
	myint := 100
	my := map[string]any{"int": myint, "str": mystr}
	debug.PtA(my)
	debug.Pt("mystr", mystr)
	debug.Pt("my", my)
}
