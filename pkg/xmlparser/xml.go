package xmlparser

import (
	"encoding/xml"
	"fmt"
	"regexp"

	"github.com/derekbelloni/go-sitemap-builder/pkg/links"
)
type URL struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	Urls []URL `xml:"urlset>url"`
}


func MatchUrls(links []links.Link) []string {
	pattern := `^(https://(www|courses)\.calhoun\.io)?/[^/]*(/.*)?$`
	re := regexp.MustCompile(pattern)
	
	var matchedItems []string

	for _, link := range links {
		if re.MatchString(link.Href) {
			matchedItems = append(matchedItems, link.Href)
		}
	}
	return matchedItems
}


func MarshalXML(SiteLinks []string) ([]byte, error) {
	urls := make([]URL, len(SiteLinks));

	for i, link := range SiteLinks {
		urls[i] = URL{Loc: link}
	}	
	sitemap := Urlset{Urls: urls}

	data, err := xml.MarshalIndent(sitemap, "", "   ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal xml: %v", err)
	}
	return data, nil
}