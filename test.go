package main

import (
	"fmt"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

type SiteConfig struct {
	Title  string
	Body   string
	Author string
	Date   string
}

type Document struct {
	Provider string
	Title    string
	Body     string
	Author   string
	Date     string
}

func main() {

	lemondeconfig := &SiteConfig{
		Title:  "//title",
		Body:   "//p[contains(@class, 'js-tweet-text')]",
		Author: "//strong[contains(@class, 'fullname')]",
		Date:   "//span[contains(@class, 'js-short-timestamp')]/@data-time",
	}
	url := "https://twitter.com/Interior/status/510085747315978241"
	document := &Document{}
	page, _ := http.Get(url)
	root, _ := xmlpath.ParseHTML(page.Body)

	document.Provider = "twitter"
	title := xmlpath.MustCompile(lemondeconfig.Title)
	document.Title, _ = title.String(root)

	body := xmlpath.MustCompile(lemondeconfig.Body)
	document.Body, _ = body.String(root)

	author := xmlpath.MustCompile(lemondeconfig.Author)
	document.Author, _ = author.String(root)

	date := xmlpath.MustCompile(lemondeconfig.Date)
	document.Date, _ = date.String(root)

	fmt.Println("Original URL", url, "\nTitle:", document.Title, "\nAuthor:", document.Author, "\nbody: ", document.Body)

}
