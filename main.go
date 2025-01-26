package main

import (
	_ "CLASS-TRACKER/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

