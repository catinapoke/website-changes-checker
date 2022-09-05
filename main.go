package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/catinapoke/website-changes-checker/additional"
	"github.com/catinapoke/website-changes-checker/basic"
)

func debug_document(document *goquery.Document, filepath string) {
	content, err := goquery.OuterHtml(document.Selection)

	if err != nil {
		log.Printf("Can't get Outer html: %v", err)
		return
	}

	err = os.WriteFile(filepath, []byte(content), 0644)

	if err != nil {
		log.Printf("Can't write file: %v", err)
		return
	}
}

func CheckPage(webpage, selector string, cookies map[string]string) {
	// var requester basic.IWebpageRequester = basic.BasicRequester{CookiesMap: cookies}
	var requester basic.IWebpageRequester = &additional.HeadlessBrowserRequester{TimeoutSeconds: 1}
	var parser basic.IContentParser = &basic.BasicParser{}
	var hasher basic.IHasher = &basic.HasherSha1{}
	var hashOperator basic.IHashOperator = &basic.BasicHashOperator{
		Url:     webpage,
		Actions: []basic.IResultAction{&basic.DebugResultAction{}},
	}

	doc, err := requester.GetPage(webpage)

	if err != nil {
		hashOperator.HandleNewHash("", err)
		return
	}

	debug_document(doc, "page.html")
	content, err := parser.GetContent(doc, selector)

	if err != nil {
		hashOperator.HandleNewHash("", err)
		return
	}

	fmt.Println(content)
	hash, err := hasher.GetHash(content)

	if err != nil {
		hashOperator.HandleNewHash("", err)
		return
	}

	hashOperator.HandleNewHash(hash, nil)
}

func main() {
	webpage := "https://www.google.com/"
	selector := ".lnXdpd"
	for {
		CheckPage(webpage, selector, nil)
		time.Sleep(1 * time.Minute)
	}
}
