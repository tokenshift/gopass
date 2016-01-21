package main

import (
	"fmt"

	// . "github.com/tokenshift/args"
	"github.com/tokenshift/gopass/lib"
)

func main() {
	fmt.Println(lib.RandomPassword(24))
}