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
	debug.PtT("Debug Usage")
	defer debug.PtT("end of Debug Usage")
	debug.PtA(my)
	debug.PtV("mystr", mystr)
	debug.PtV("my", my)
}
