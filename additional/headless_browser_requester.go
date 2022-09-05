package additional

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/playwright-community/playwright-go"
)

type HeadlessBrowserRequester struct {
	TimeoutSeconds int
}

func (x *HeadlessBrowserRequester) GetPage(url string) (*goquery.Document, error) {
	pw, err := playwright.Run()

	if err != nil {
		log.Printf("Playwright couldn't be started: %v", err)
		return nil, err
	}

	browser, err := pw.Firefox.Launch()

	if err != nil {
		log.Printf("Playwright couldn't start browser: %v", err)
		return nil, err
	}

	page, err := browser.NewPage()

	if err != nil {
		log.Printf("Playwright browser couldn't create page: %v", err)
		return nil, err
	}

	_, err = page.Goto(url)

	if err != nil {
		log.Printf("Playwright browser couldn't load page: %v", err)
		return nil, err
	}

	time.Sleep(1 * time.Second)

	html, err := page.Content()

	if err != nil {
		log.Printf("Playwright browser couldn't get page content: %v", err)
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Printf("Couldn't convert html to document: %v", err)
		return nil, err
	}

	return document, nil
}
