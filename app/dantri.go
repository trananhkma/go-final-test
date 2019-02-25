package crawler

type DanTri struct{}

func (d *DanTri) GetTitle(doc interface{}) string {
	return doc.Find(".details__headline").Text()
}

func (d *DanTri) GetDescription(doc interface{}) string {
	return doc.Find(".sapo").Text()
}

func (d *DanTri) GetBody(doc interface{}) string {
	body, _ := doc.Find(".cms-body").Html()
	return body
}

func (d *DanTri) GetTime(doc interface{}) string {
	return doc.Find(".details__meta .meta time").Text()
}
