package main

import (
	sps "lit-news/spiders"
	utils "lit-news/utils"
)

func main() {
	utils.DbAdd(sps.Jwc())
	utils.DbAdd(sps.Xwzx())
	utils.DbAdd(sps.Tw())
	defer utils.DbClose()
}
