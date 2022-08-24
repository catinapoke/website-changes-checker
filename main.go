package main

import (
	"time"

	"github.com/catinapoke/website-changes-checker/basic"
)

func CheckPage(webpage, selector string, cookies map[string]string) {
	var requester basic.IWebpageRequester = basic.BasicRequester{CookiesMap: cookies}
	var parser basic.IContentParser = basic.BasicParser{}
	var hasher basic.IHasher = basic.HasherSha1{}
	var hashOperator basic.IHashOperator = basic.BasicHashOperator{
		Url:     webpage,
		Actions: []basic.IResultAction{basic.DebugResultAction{}},
	}

	doc := requester.GetPage(webpage)
	content := parser.GetContent(doc, selector)
	hash := hasher.GetHash(content)
	hashOperator.HandleNewHash(hash)
}

func main() {
	webpage := "https://www.google.com/"
	selector := ".lnXdpd"
	for {
		CheckPage(webpage, selector, nil)
		time.Sleep(5 * time.Minute)
	}
}
