package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Db 数据库连接
var Db *gorm.DB

//Newrow 单个新闻模板
type Newrow struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Date    string `json:"date"`
	Type    string `json:"type"`
	SubType string `json:"subtype"`
}

//New 新闻返回模板
type New struct {
	gorm.Model
	Newrow
}

func init() {
	db, err := gorm.Open("mysql", "root:root@/lit?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&New{})
	if err != nil {
		println(err)
	}
	Db = db
}

//DbClose 关闭数据库连接
func DbClose() {
	Db.Close()
}

//DbAdd 增加数据
func DbAdd(allData []map[string]string) {
	for _, value := range allData {
		news := New{}
		Db.Where("link = ?", value["link"]).Find(&news)
		lnews := New{Newrow: Newrow{Title: value["title"], Link: value["link"], Date: value["date"], Type: value["type"], SubType: value["subtype"]}}
		if news.ID != 0 {
			Db.First(&lnews)
			fmt.Print(value["title"] + " 存在！更新一下\n")
			Db.Save(&lnews)
		} else {
			Db.Create(&lnews)
			fmt.Print("新增 " + value["title"] + "\n")
		}
	}
}

//DbGetNews 根据分类获取新闻
func DbGetNews(t string) (newrow []Newrow) {

	if t == "new" {
		Db.Table("news").Select([]string{"title", "link", "date", "type", "sub_type"}).Order("date desc").Limit(20).Find(&newrow)
	} else {
		Db.Table("news").Select([]string{"title", "link", "date", "type", "sub_type"}).Where("type = ?", t).Order("date desc").Limit(15).Find(&newrow)
	}

	return
}

//DbGetUpdate 获取更新时间
func DbGetUpdate() (update gorm.Model) {
	Db.Table("news").Select("updated_at").Order("updated_at desc").First(&update)
	return
}
