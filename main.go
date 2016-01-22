package main

import (
	"fmt"

	// . "github.com/tokenshift/args"
	"github.com/tokenshift/gopass/lib"
)

func main() {
	fmt.Println(lib.RandomPasswordOpts(24, lib.DEFAULT_OPTIONS))
	fmt.Println(lib.RandomPasswordOpts(24, lib.DEFAULT_OPTIONS ^ lib.INCLUDE_SPECIAL))
	fmt.Println(lib.XKCDPassword(24))
}