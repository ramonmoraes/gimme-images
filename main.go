package main

import (
	"fmt"
	"io/ioutil"
	"log"

	gimme "ramonmoraes/gimme-images/gimme"
)

func main() {
	crawler := gimme.Crawler{Domain: "https://ocapacitor.com"}
	urls := crawler.CrawlURL("https://ocapacitor.com/john-wick-3-parabellum-filme-tem-novas-cenas-divulgadas/")
	saveSrcs(urls)
}

func saveSrcs(srcs []string) {
	errors := []error{}
	for _, src := range srcs {
		img := gimme.Image{URL: src}
		d_err := img.Download()
		if d_err != nil {
			errors = append(errors, d_err)
		}
		s_err := img.Save()
		if s_err != nil {
			errors = append(errors, s_err)
		}
	}
	logErrors(errors)
}

func logErrors(errors []error) {
	for _, e := range errors {
		fmt.Println(e)
	}
}

func sample() {
	bytes, err := ioutil.ReadFile("data/fixture.html")
	if err != nil {
		log.Fatal(err)
	}

	doc := gimme.GetDocument(bytes)
	srcs := gimme.GetImagesSRC(doc)
	saveSrcs(srcs)
}
