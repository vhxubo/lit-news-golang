package spiders

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Xwzx 洛阳理工学院新闻中心首页新闻
func Xwzx() []map[string]string {
	Type := "xwzx"
	subTypes := []string{
		"xwkx",
		"ggtz",
		"xsxx",
		"mtxw",
	}
	var allData []map[string]string
	host := "http://www.lit.edu.cn/xwzx/"

	for _, subType := range subTypes {
		url := host + subType + ".htm"
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(url + "爬取失败")
		} else {
			fmt.Println("开始爬取: " + url)
			doc, _ := goquery.NewDocumentFromReader(res.Body)
			doc.Find(".wp_article_list").Find("li").Each(func(i int, s *goquery.Selection) {
				title, _ := s.Find("a").Attr("title")
				link, _ := s.Find("a").Attr("href")
				date := s.Find(".Article_PublishDate").Text()
				dateLen := len(date)
				if dateLen > 18 {
					date = date[dateLen-18 : dateLen]
				}
				date = strings.ReplaceAll(date, "年", "-")
				date = strings.ReplaceAll(date, "月", "-")
				date = strings.ReplaceAll(date, "日", "")
				date = strings.Trim(date, " ")

				allData = append(allData, map[string]string{"title": title, "link": host + link, "date": date, "type": Type, "subtype": subType})

			})
		}
		defer res.Body.Close()
	}
	return allData
}
