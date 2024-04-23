package main

import (
	_ "image/png"
	"log"

	"github.com/eric-lab-star/ebitGame/res/run"
)

func main() {
	run.InitImgs()
	is := run.Imgs
	log.Printf("is: %v\n", is)
}
