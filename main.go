package main

import (
	_ "beegoapidemo/routers"
	"github.com/astaxie/beego"
	_ "beegoapidemo/base"
)

func main() {
	beego.Run()
}
