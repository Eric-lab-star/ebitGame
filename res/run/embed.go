package run

import (
	"bytes"
	"embed"
	"log"
)

//go:embed Run01.png Run02.png Run03.png Run03.png Run04.png Run05.png Run06.png
var Res embed.FS
var Imgs []*bytes.Reader

func InitImgs() {
	dirEntries, err := Res.ReadDir(".")
	if err != nil {
		log.Printf("err: %v\n", err)
	}

	for _, v := range dirEntries {
		b, err := Res.ReadFile(v.Name())
		if err != nil {
			log.Printf("err %v\n", err)
		}
		br := bytes.NewReader(b)
		Imgs = append(Imgs, br)

	}
}
