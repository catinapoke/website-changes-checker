package basic

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type IWebpageRequester interface {
	GetPage(url string) *goquery.Document
}

type BasicRequester struct {
	CookiesMap map[string]string
}

func (requester BasicRequester) SetCookies(CookiesMap map[string]string) {
	requester.CookiesMap = CookiesMap
}

func (requester BasicRequester) GetPage(url string) *goquery.Document {
	var response *http.Response
	var err error

	if len(requester.CookiesMap) == 0 {
		response, err = http.Get(url)
	} else {
		response, err = requester.getRequest(url)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Couldn't fetch data: %d %s\n", response.StatusCode, response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return document
}

func (requester BasicRequester) getRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range requester.CookiesMap {
		request.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	client := &http.Client{}
	return client.Do(request)
}
