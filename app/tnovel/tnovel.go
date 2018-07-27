package main

import (
	"github.com/jrecuero/go-cli/app/tnovel/novel"
	"github.com/jrecuero/go-cli/tools"
)

func main() {
	//defer func() {
	//    if r := recover(); r != nil {
	//        tools.ToDisplay("Error: %#v\n", r)
	//    }
	//}()
	_novel := &novel.Novel{}
	aeStrings := []string{
		"A1 hit T1",
		"A1 heal <self>",
		"A1 skil [T1 T2]",
		"A1 skill] T1",
		"A1 skill T1",
	}
	for _, str := range aeStrings {
		ae := _novel.Compile(str)
		if ae != nil {
			tools.ToDisplay("%#v\n", ae)
		} else {
			tools.ToDisplay("--- INVALID ACTION ---\n")

		}
	}
}
