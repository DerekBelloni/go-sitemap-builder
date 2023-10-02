package xmlparser

import (
	"fmt"
	"encoding/xml"
	"github.com/derekbelloni/go-sitemap-builder/pkg/links"
)
type URL struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	Urls []URL `xml:"urlset>url"`
}
func MarshalXML(SiteLinks []links.Link) ([]byte, error) {
	urls := make([]URL, len(SiteLinks));

	for i, link := range SiteLinks {
		urls[i] = URL{Loc: link.Href}
	}	
	sitemap := Urlset{Urls: urls}

	data, err := xml.MarshalIndent(sitemap, "", "   ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal xml: %v", err)
	}
	return data, nil
}