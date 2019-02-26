package crawler

import (
    "github.com/PuerkitoBio/goquery"
    "strings"
)

type ThanhNien struct{}

func (t *ThanhNien) GetTitle(doc *goquery.Document) string {
	return doc.Find(".details__headline").Text()
}

func (t *ThanhNien) GetDescription(doc *goquery.Document) string {
	return doc.Find(".sapo").Text()
}

func (t *ThanhNien) GetBody(doc *goquery.Document) string {
	body, _ := doc.Find(".cms-body").Html()
	return body
}

func (th *ThanhNien) GetTime(doc *goquery.Document) string {
    time := doc.Find(".details__meta .meta time").Text()
	t := strings.Split(time, " - ")
	return t[1] + " " + t[0]
}
