package crawler

type ThanhNien struct{}

func (t *ThanhNien) GetTitle(doc interface{}) string {
	return doc.Find(".details__headline").Text()
}

func (t *ThanhNien) GetDescription(doc interface{}) string {
	return doc.Find(".sapo").Text()
}

func (t *ThanhNien) GetBody(doc interface{}) string {
	body, _ := doc.Find(".cms-body").Html()
	return body
}

func (th *ThanhNien) GetTime(doc interface{}) string {
	time = doc.Find(".details__meta .meta time").Text()
	t := strings.Split(time, " - ")
	return t[1] + " " + t[0]
}
