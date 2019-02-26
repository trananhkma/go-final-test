package crawler

import "github.com/PuerkitoBio/goquery"

type DanTri struct{}

func (d *DanTri) GetTitle(doc *goquery.Document) string {
	return doc.Find(".details__headline").Text()
}

func (d *DanTri) GetDescription(doc *goquery.Document) string {
	return doc.Find(".sapo").Text()
}

func (d *DanTri) GetBody(doc *goquery.Document) string {
	body, _ := doc.Find(".cms-body").Html()
	return body
}

func (d *DanTri) GetTime(doc *goquery.Document) string {
	return doc.Find(".details__meta .meta time").Text()
}
