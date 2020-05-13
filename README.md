# lit-news-golang

需要配合 [lit-news-web](https://github.com/vhxubo/lit-news-web) 使用

## 获取新闻

MySQL 需要新建名为 lit 的数据库，账号密码均为 `root`

运行 `GetNews.go`，spiders 目录为各个网页对应的爬虫

## 服务端

运行 `Server.go`，访问 API 接口

## 使用

依次编译 `GetNews.go`、`GetNews.go`，在服务器上执行 `cron.sh` 即可

API:

- 最新: http://localhost:8080/getnews?type=new
- 教务处: http://localhost:8080/getnews?type=jwc
- 新闻中心: http://localhost:8080/getnews?type=xwzx
- 团委: http://localhost:8080/getnews?type=tw
