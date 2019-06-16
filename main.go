package main

import (
	"io/ioutil"
	"log"
	gimme "ramonmoraes/gimme-images/gimme"
)

func main() {
	urls := gimme.CrawlURL("https://ocapacitor.com/john-wick-3-parabellum-filme-tem-novas-cenas-divulgadas/")
	errs := gimme.DownloadURLS(urls)
	if len(errs) != 1 {
		log.Fatal(errs)
	}
}

func sample() {
	bytes, err := ioutil.ReadFile("data/fixture.html")
	if err != nil {
		log.Fatal(err)
	}

	doc := gimme.GetDocument(bytes)
	srcs := gimme.GetImagesSRC(doc)
	for _, src := range srcs {
		img := gimme.Image{URL: src}
		img.Download()
		img.Save()
	}
}
