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
	url := flag.String("url", "https://gophercises.com", "supply a url")

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

	domainLinks := xmlparser.MatchUrls(links)

	for _, item := range domainLinks {
		fmt.Printf("item:%v\n",item)
	}

	xmlData, err := xmlparser.MarshalXML(domainLinks)

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

 
