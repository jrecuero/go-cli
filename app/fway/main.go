package main

import (
	"fmt"

	"github.com/jrecuero/go-cli/apps/freeway"
)

func main() {
	fway := freeway.NewFreeway()
	fmt.Println(fway.String())
}
