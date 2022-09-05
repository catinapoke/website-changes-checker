package basic

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type IWebpageRequester interface {
	GetPage(url string) (*goquery.Document, error)
}

type BasicRequester struct {
	CookiesMap map[string]string
}

func (requester BasicRequester) SetCookies(CookiesMap map[string]string) {
	requester.CookiesMap = CookiesMap
}

func (requester BasicRequester) GetPage(url string) (*goquery.Document, error) {
	var response *http.Response
	var err error

	if len(requester.CookiesMap) == 0 {
		response, err = http.Get(url)
	} else {
		response, err = requester.getRequest(url)
	}

	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Couldn't fetch data: %d %s\n", response.StatusCode, response.Status))
		log.Print(err)
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return document, nil
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
