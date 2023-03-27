package main

import (
	"NotePad/conf"
	"NotePad/route"
)

func main() {
	// 配置初始化
	conf.Init()
	// 路由初始化
	r := route.NewRouter()
	//运行
	r.Run(conf.SERVER_PORT)
}
