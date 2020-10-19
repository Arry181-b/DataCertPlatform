package main

import (
	"BCP/db_mysql"
	_ "BCP/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()

	beego.SetStaticPath("/js","./static.js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

