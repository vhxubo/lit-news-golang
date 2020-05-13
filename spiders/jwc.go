package spiders

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Jwc 洛阳理工学院教务处首页新闻
func Jwc() []map[string]string {

	Type := "jwc"
	var allData []map[string]string
	host := "http://www.lit.edu.cn/jwc/"

	res, err := http.Get(host)
	if err != nil {
		fmt.Println(host + "爬取失败")
	} else {
		fmt.Println("开始爬取: " + host)
		doc, _ := goquery.NewDocumentFromReader(res.Body)
		doc.Find(".post_box").Each(func(i int, s *goquery.Selection) {
			date, _ := s.Find(".tit").Prev().Html()
			date = strings.Trim(date, " \n\t\r")
			date = strings.ReplaceAll(date, "<br/>", "-")
			title := s.Find("h2 > a").Text()
			link, _ := s.Find("h2 > a").Attr("href")
			allData = append(allData, map[string]string{"title": title, "link": host + link, "date": date, "type": Type, "subtype": ""})
		})
	}
	defer res.Body.Close()
	return allData
}
