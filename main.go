package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-final-test/app"
	"net/http"
	"strings"
)

type Target interface {
	GetTitle() string
	GetDescription() string
	GetBody() string
	GetTime() string
}

func getTarget(url string) Target {
	var target struct{}
	if strings.Contains(url, "https://thanhnien.vn") {
		target = crawler.ThanhNien{}
	} else {
		target = crawler.DanTri{}
	}

	return target
}

func Crawler() {
	http.HandleFunc("/crawler", func(response http.ResponseWriter, request *http.Request) {
		//url := "https://thanhnien.vn/thoi-su/hinh-anh-nguoi-dan-ong-dung-xe-de-di-ve-sinh-giua-duong-cao-toc-gay-phan-no-1037558.html"
		url, _ := request.URL.Query()["url"]
		resp, e := http.Get(url[0])

		if e != nil {
			panic(e)
		}
		defer resp.Body.Close()

		target := getTarget(url[0])

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			panic(err)
		}

		title := target.GetTitle(doc)
		fmt.Fprintln(response, "<html><body><h3>Title: ", title, "</h3><p></p>")

		des := target.GetDescription(doc)
		fmt.Fprintln(response, "<h3>Desciption: ", des, "</h3><p></p>")

		body := target.GetBody(doc)
		fmt.Fprintln(response, "<h3>Body: </h3><p></p>")
		fmt.Fprintln(response, body)

		time := target.GetTime(doc)
		fmt.Fprintln(response, "<h3>Time: ", time, "</h3><p></p>")

		fmt.Fprintln(response, "</body></html>")
	})

	http.ListenAndServe("localhost:6969", nil)
}

func main() {
	Crawler()
}
