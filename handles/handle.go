package handles

import (
	"github.com/Tuan8888/falcon-callback-plugin/configs"
	"github.com/Tuan8888/falcon-callback-plugin/logs"
	"net/http"
	"os/exec"
)

func CommandHandle(w http.ResponseWriter, r *http.Request) {
	logs.Info.Println("接收到地址", r.RemoteAddr, "发送的请求", r.URL)
	r.ParseForm()
	key := r.Form["key"]
	if key == nil {
		logs.Error.Println("发送请求中参数 key 不正确")
		w.Write([]byte("发送请求中参数 key 不正确"))
		return
	} else {
		command := configs.Maps[key[0]]
		if command == "" {
			logs.Error.Println(key[0], "对应的远程命令未配置")
		} else {
			logs.Info.Println("查询到远程命令", command)
			w.Write([]byte(execShell(command)))
		}
	}
}

/*
	执行shell命令
*/
func execShell(command string) []byte {
	logs.Info.Println("准备执行命令", command)
	cmd := exec.Command("/bin/bash -c", "'"+command+"'")
	res, err := cmd.Output()
	if err != nil {
		logs.Error.Println("执行命令出错", err)
		return []byte(err.Error())
	}
	logs.Info.Println("命令执行完毕，返回")
	logs.Info.Println(string(res))
	return res
}
