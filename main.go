package main

import (
	"fmt"

	"github.com/rafaelpapastamatiou/go-lang-object-oriented/arrays"
	"github.com/rafaelpapastamatiou/go-lang-object-oriented/interfaces"
	"github.com/rafaelpapastamatiou/go-lang-object-oriented/maps"
	"github.com/rafaelpapastamatiou/go-lang-object-oriented/slices"
	"github.com/rafaelpapastamatiou/go-lang-object-oriented/structs"
)

func main() {
	arrays.Describe()
	separator()
	slices.Describe()
	separator()
	maps.Describe()
	separator()
	structs.Describe()
	separator()
	interfaces.Describe()
}

func separator() {
	fmt.Printf("\n\n\n#===================================================================#\n\n\n")
}
