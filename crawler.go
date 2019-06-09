package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getBodyFromURL(URL string) []byte {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getDocument(body []byte) *goquery.Document {
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func getImagesSRC(doc *goquery.Document) []string {
	urls := []string{}
	doc.Find("img").Each(func(i int, sel *goquery.Selection) {
		urls = append(urls, sel.AttrOr("src", ""))
	})
	return urls
}
