package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/eric-lab-star/ebitGame/res/run"
)

func main() {
	b := run.Run_png
	img, name, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	fmt.Println(img)
}
