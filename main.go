package main

import (
	_ "bet365/bootstrap"
	"bet365/routers"
	"bet365/weChat/global/variable"
)

// 这里可以存放门户类网站入口
func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
