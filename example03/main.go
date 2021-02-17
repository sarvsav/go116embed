package main

import (
	"embed"
	"fmt"
)

//go:embed dist/*
var f embed.FS // To read from file system

func main() {

	dataFileA, _ := f.ReadFile("dist/a.txt")
	fmt.Println(string(dataFileA))
	dataFileB, _ := f.ReadFile("dist/b.txt")
	fmt.Println(string(dataFileB))
}
