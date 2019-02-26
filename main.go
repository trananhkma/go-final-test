package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type Info struct {
	title       string
	description string
	body        string
	time        string
}

type DataCrawler interface {
	CrawlData(doc *goquery.Document) Info
}

type ThanhNienCrawler struct{}
func (th ThanhNienCrawler) CrawlData(doc *goquery.Document) Info {
	var info Info
	info.title = doc.Find(".details__headline").Text()
	info.description = doc.Find(".sapo").Text()
	info.body, _ = doc.Find(".cms-body").Html()
	time := doc.Find(".details__meta .meta time").Text()
	t := strings.Split(time, " - ")
	info.time = t[1] + " " + t[0]
	return info
}

type DanTriCrawler struct{}
func (d DanTriCrawler) CrawlData(doc *goquery.Document) Info {
	var info Info
	info.title = doc.Find(".fon31").Text()
	info.description = doc.Find(".fon33").Text()
	info.body = doc.Find(".fon34").Text()
	info.time = doc.Find(".fon7").Text()
	return info
}


func Crawl(domain string, doc *goquery.Document) Info {
	var crawler DataCrawler
	switch domain {
	case "thanhnien":
		crawler = ThanhNienCrawler{}
		return crawler.CrawlData(doc)
	case "dantri":
		crawler = DanTriCrawler{}
		return crawler.CrawlData(doc)
	}
	return Info{}
}

func getDomain(url string) string {
	s := strings.Split(url, "://")
	uri := s[1]
	s = strings.Split(uri, ".")
	return s[0]
}

func CrawlServer() {
	http.HandleFunc("/crawler", func(response http.ResponseWriter, request *http.Request) {
		url, _ := request.URL.Query()["url"]
		resp, e := http.Get(url[0])

		if e != nil {
			panic(e)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}

		domain := getDomain(url[0])
		info := Crawl(domain, doc)

		fmt.Fprintln(response, "<html><body><h3>Title: ", info.title, "</h3><p></p>")
		fmt.Fprintln(response, "<h3>Desciption: ", info.description, "</h3><p></p>")
		fmt.Fprintln(response, "<h3>Body: </h3><p></p>")
		fmt.Fprintln(response, info.body)
		fmt.Fprintln(response, "<h3>Time: ", info.time, "</h3><p></p>")
		fmt.Fprintln(response, "</body></html>")
	})
	http.ListenAndServe("localhost:6969", nil)
}

func main() {
	CrawlServer()
}
