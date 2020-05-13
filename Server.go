package main

import (
	utils "lit-news/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CorsMiddleware())
	r.GET("/getnews", func(c *gin.Context) {
		newsType := c.DefaultQuery("type", "new")
		var lists []utils.Newrow

		if newsType == "new" || newsType == "jwc" || newsType == "xwzx" || newsType == "tw" {
			lists = utils.DbGetNews(newsType)
			update := utils.DbGetUpdate()
			c.JSON(200, gin.H{"code": 200, "update": update.UpdatedAt.Format("2006-01-02 15:04:05"), "lists": lists})
		} else {
			c.JSON(406, gin.H{"code": 406, "msg": "error"})
		}

	})
	r.Run()
}

// CorsMiddleware dsfsd
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}

}
