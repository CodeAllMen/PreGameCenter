package main

import (
	"github.com/MobileCPX/PreGameCenter/initial"

	_ "github.com/MobileCPX/PreGameCenter/routers"
	"github.com/astaxie/beego"
)

func main() {
	initial.FileToSql()

	beego.Run()
}
