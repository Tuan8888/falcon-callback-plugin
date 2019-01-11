package callback_plugin

import (
	"callback-plugin/configs"
	"callback-plugin/logs"
	"callback-plugin/routes"
)

func main() {
	// 初始化日志记录器
	logs.Init()

	// 加载配置文件
	configs.Init()

	// 启动 API 服务
	routes.Init()
}
