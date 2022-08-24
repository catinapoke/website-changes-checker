package basic

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

type IContentParser interface {
	GetContent(document *goquery.Document, css_selector string) string
}

type BasicParser struct{}

func (BasicParser) GetContent(document *goquery.Document, css_selector string) string {
	content := document.Find(css_selector)
	html, err := content.First().Html()

	if err != nil {
		log.Fatal(err)
	}

	return html
}
