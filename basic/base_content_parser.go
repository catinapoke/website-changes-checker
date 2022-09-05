package basic

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

type IContentParser interface {
	GetContent(document *goquery.Document, css_selector string) (string, error)
}

type BasicParser struct{}

func (BasicParser) GetContent(document *goquery.Document, css_selector string) (string, error) {
	content := document.Find(css_selector)
	html, err := content.First().Html()

	if err != nil {
		log.Print(err)
		return "", err
	}

	return html, nil
}
