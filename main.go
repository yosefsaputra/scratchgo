package main

import (
	"github.com/yosefsaputra/scratchgo/debug"
)

func main() {
	mystr := "test"
	myint := 100
	myvarmap := map[string]any{"int": myint, "str": mystr}
	mylist := []string{"string1", "string2"}
	myperson := Person{Name: "person1", Age: 12}
	mypersonlist := []Person{{Name: "person1", Age: 12}, {Name: "person2", Age: 13}}
	mygroup := Group{Name: "group1", Members: mypersonlist}

	// Usage
	debug.PtT("Debug Usage")
	defer debug.PtT("end of Debug Usage")
	debug.Ptln()

	debug.Pt("Print Any")
	debug.PtA(mystr)
	debug.PtA(myint)
	debug.PtA(myvarmap)
	debug.PtA(mylist)
	debug.PtA(myperson)
	debug.PtA(mypersonlist)
	debug.PtA(&myperson)
	debug.PtA(mygroup)
	debug.Ptln()

	debug.Pt("Print Variable")
	debug.PtV("mystr", mystr)
	debug.PtV("myint", myint)
	debug.PtV("myvarmap", myvarmap)
	debug.PtV("mylist", mylist)
	debug.PtV("myperson", myperson)
	debug.PtV("mypersonlist", mypersonlist)
	debug.PtV("mypersonpointer", &myperson)
	debug.PtV("mygroup", mygroup)
	debug.Ptln()
}

type Person struct {
	Name string
	Age  int
}

type Group struct {
	Name    string
	Members []Person
}
