package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"
	"os"
)

//go:embed gopher.png
var images []byte

func generateImage(imageByte []byte) {

	img, _, err := image.Decode(bytes.NewReader(imageByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./result.png")
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		log.Println(err)
	}

}

func main() {
	generateImage(images)
}
