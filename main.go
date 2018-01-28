package main

import (
	_ "github.com/astaxie/beego"
	_ "github.com/neatLines/beeDemo/routers"
)

func main() {
	beego.Run()
}
