package configs

import (
	"callback-plugin/logs"
	"encoding/json"
	"os"
)

var Maps = make(map[string]string)	//get 请求参数，对应的远程命令
var Cfg Config										// 配置信息

type Config struct {
	Host			string						`json:"host"`
	Port				string						`json:"port"`
	Mappers		[]Mapper					`json:"mappers"`
}

type Mapper struct {
	Key					string		`json:"key"`
	Command		string		`json:"command"`
}

func Init() {
	file, err := os.Open("cfg.json")

	if err != nil {
		logs.Error.Println("没有找到配置文件cfg.json，请检查")
		os.Exit(1)
	}

	err = json.NewDecoder(file).Decode(&Cfg)
	if err != nil {
		logs.Error.Println("cfg.json解析错误")
		os.Exit(2)
	}

	for _, mapper := range Cfg.Mappers {
		Maps[mapper.Key] = mapper.Command
	}

	logs.Info.Println("读取配置文件完毕")
	logs.Info.Println("Host的值为", Cfg.Host)
	logs.Info.Println("启动服务端口为", Cfg.Port)
	logs.Info.Println("配置映射如下")


	for key, command := range Maps {
		logs.Info.Printf("参数 %s 对应的远程命令为 %s", key, command)
	}
}
