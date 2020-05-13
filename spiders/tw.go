package spiders

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//Tw 洛阳理工学院团委首页新闻
func Tw() []map[string]string {

	Type := "tw"
	subTypes := []string{
		"qnkx",
		"tntz",
	}
	var allData []map[string]string
	host := "http://www.lit.edu.cn/tw/"

	for _, subType := range subTypes {
		url := host + subType + ".htm"
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(url + "爬取失败")
		} else {
			fmt.Println("开始爬取: " + url)
			doc, _ := goquery.NewDocumentFromReader(res.Body)
			doc.Find(".caselist").Find("li").Each(func(i int, s *goquery.Selection) {
				title := s.Find("a").Text()
				link, _ := s.Find("a").Attr("href")
				date := s.Find(".time").Text()
				dateLen := len(date)
				if dateLen > 10 {
					date = date[:10]
				}
				allData = append(allData, map[string]string{"title": title, "link": host + link, "date": date, "type": Type, "subtype": subType})
			})
		}
		defer res.Body.Close()
	}
	return allData
}
