package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile("./fixture.html")
	if err != nil {
		log.Fatal(err)
	}

	doc := getDocument(bytes)
	imgs := getImagesSRC(doc)
	fmt.Println(imgs)
}

func example() {
	url := "https://i2.wp.com/ocapacitor.com/wp-content/uploads/2016/08/0000jgpfymtl.jpg?fit=655%2C370&ssl=1&resize=350%2C200"
	filepath := ".dart/img.jpg"
	err := downloadImag(filepath, url)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadImag(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
