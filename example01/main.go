package main

import (
	_ "embed"

	"fmt"
)

//go:embed config
var s string

func main() {
	fmt.Println(s)
}
