crontab -e 
#
0 7-23/2 * * * /www/wwwroot/lit-news/GetNews >> /www/wwwroot/lit-news/GetNews.log
nohup /www/wwwroot/lit-news/Server &
