package gimme

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func DownloadURLS(urls []string) []error {
	errors := []error{}
	for _, url := range urls {
		img := Image{URL: url}
		err := img.Download()
		if err != nil {
			errors = append(errors, err)
		}
		err = img.Save()
	}
	return errors
}

func CrawlURL(URL string) []string {
	body := GetBodyFromURL(URL)
	doc := GetDocument(body)
	return GetImagesSRC(doc)
}

func GetBodyFromURL(URL string) []byte {
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

func GetDocument(body []byte) *goquery.Document {
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetImagesSRC(doc *goquery.Document) []string {
	urls := []string{}
	doc.Find("img").Each(func(i int, sel *goquery.Selection) {
		urls = append(urls, sel.AttrOr("src", ""))
	})
	return urls
}
