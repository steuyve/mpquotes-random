package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://mathprofessorquotes.com/random")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalln("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(body.Find(".quote").Text())
	fmt.Println(body.Find(".cont.quote_source").Text())

	// If the post is in conversation format
	body.Find(".chat_line").Each(func(index int, item *goquery.Selection) {
		speaker := strings.TrimSpace(item.RemoveFiltered("strong").Text())
		fmt.Printf("%s %s\n", strings.TrimSpace(speaker), strings.TrimSpace(item.Text()))
	})

	fmt.Printf("\nLink: %s\n", res.Request.URL.String())
}
