package main

import (
	"github.com/Tuan8888/falcon-callback-plugin/configs"
	"github.com/Tuan8888/falcon-callback-plugin/logs"
	"github.com/Tuan8888/falcon-callback-plugin/routes"
)

func main() {
	// 初始化日志记录器
	logs.Init()

	// 加载配置文件
	configs.Init()

	// 启动 API 服务
	routes.Init()
}
