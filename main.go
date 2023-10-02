package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"github.com/derekbelloni/go-sitemap-builder/pkg/links"
	"github.com/derekbelloni/go-sitemap-builder/pkg/xmlparser"
)

func main() {
	url := flag.String("url", "https://go.dev/doc/", "supply a url")

	flag.Parse();

	body, err := fetchHTML(*url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	r := strings.NewReader(body)

	links, err := links.Parse(r)
	if err != nil {
		log.Printf("Error retrieving links from HTML: %v", err)
	}

	xmlData, err := xmlparser.MarshalXML(links)

	if err != nil {
		log.Printf("Error retrieving xml sitemap: %v", err)
	}

	fmt.Println(string(xmlData))
}

func fetchHTML(url string) (string, error) {
	resp, err:= http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error reading the body for URL %s: %w", url, err)
	}
	return string(body), nil
 }
