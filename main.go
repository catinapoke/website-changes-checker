package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"crypto/sha1"
	"encoding/hex"

	"github.com/PuerkitoBio/goquery"
)

func GetRequest(url string, CookiesMap map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range CookiesMap {
		request.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	client := &http.Client{}
	return client.Do(request)
}

func get_page(page_name string, cookies map[string]string) *goquery.Document {
	var response *http.Response
	var err error

	if len(cookies) == 0 {
		response, err = http.Get(page_name)
	} else {
		response, err = GetRequest(page_name, cookies)
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

func get_content(document *goquery.Document, css_selector string) string {
	content := document.Find(css_selector)
	html, err := content.First().Html()

	if err != nil {
		log.Fatal(err)
	}

	return html
}

func get_hash(content string) string {
	hasher := sha1.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

func get_filename_by_pagename(pagename string) string {
	filename := pagename + ".txt"
	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, ":", "_")
	return filename
}

func get_saved_hash(page_name string) string {
	filename := get_filename_by_pagename(page_name)

	if !isFileExists(filename) {
		return ""
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func save_hash(hash, pagename string) {
	filename := get_filename_by_pagename(pagename)
	err := os.WriteFile(filename, []byte(hash), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func check_page(webpage, selector string, cookies map[string]string) {
	doc := get_page(webpage, cookies)
	content := get_content(doc, selector)
	hash := get_hash(content)
	saved_hash := get_saved_hash(webpage)
	if hash != saved_hash {
		// TODO: Add actions
		save_hash(hash, webpage)

		if saved_hash == "" {
			fmt.Printf("Saved initial hash \"%s\" for %s\n", hash, webpage)
		} else {
			fmt.Printf("There is changes at %s\n", webpage)
		}

	} else {
		fmt.Printf("There is no changes at %s\n", webpage)
	}
}

func main() {
	webpage := "https://www.google.com/"
	selector := ".lnXdpd"
	for {
		check_page(webpage, selector, nil)
		time.Sleep(5 * time.Minute)
	}
}
