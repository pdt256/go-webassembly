package main

import (
	"fmt"
	"go-webassembly/pkg/jdoc"
	"io/ioutil"
	"log"
	"net/http"
)

var doc = jdoc.New()

func main() {
	fmt.Println("Hello, WebAssembly console!")
	doc.Print("Hello, WebAssembly document!")

	TestHTTPRequest()
}

func TestHTTPRequest() {
	url := "https://httpbin.org/anything"
	doc.Printf("Request: %s", url)
	body, err := getBody(url)
	if err != nil {
		log.Printf("unable to get IP address: %v", err)
	} else {
		doc.Printf("Response: %s", body)
	}
}

func getBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to get request: %v", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read body: %v", err)
	}

	return string(bytes), nil
}
