package gimme

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

const BASE_PATH = "data/"

// Image is the base abstraction, it should have a content if downloaded and create it's name based on the given URL
type Image struct {
	URL     string
	Name    string
	Content []byte
}

// Download should set the image's content from the given http response
func (i *Image) Download() error {
	fmt.Printf("[Downloading]\n%s\n", i.URL)
	resp, err := http.Get(i.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	i.Content = body
	return nil
}

// Save should store the image's content on local file
func (i *Image) Save() error {
	if i.Name == "" {
		i.createName()
	}

	imgPath := fmt.Sprintf("%s%s.jpeg", BASE_PATH, i.Name)
	fmt.Printf("Saving at: %s\n", imgPath)
	out, err := os.Create(imgPath)
	if err != nil {
		return err
	}

	defer out.Close()
	out.Write(i.Content)
	return nil
}

func (i *Image) createName() {
	regex, err := regexp.Compile(".*\\/(.*)")
	if err != nil {
		log.Fatal(err)
	}
	matches := regex.FindAllStringSubmatch(i.URL, -1)
	i.Name = matches[0][len(matches[0])-1]
}
