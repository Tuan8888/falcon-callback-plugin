package routes

import (
	"github.com/Tuan8888/falcon-callback-plugin/configs"
	"github.com/Tuan8888/falcon-callback-plugin/handles"
	"net/http"
)

/*
	API
*/
func Init() {
	http.HandleFunc("/command", handles.CommandHandle)
	http.ListenAndServe(configs.Cfg.Host+":"+configs.Cfg.Port, nil)
}
