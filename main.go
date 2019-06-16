package main

import (
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("data/fixture.html")
	if err != nil {
		log.Fatal(err)
	}

	doc := getDocument(bytes)
	srcs := getImagesSRC(doc)
	for _, src := range srcs {
		img := Image{URL: src}
		img.Download()
		img.Save()
	}
}
