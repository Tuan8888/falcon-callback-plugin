package routes

import (
	"callback-plugin/configs"
	"callback-plugin/handles"
	"net/http"
)

/*
	API
 */
func Init() {
	http.HandleFunc("/command", handles.CommandHandle)
	http.ListenAndServe(configs.Cfg.Host + ":" + configs.Cfg.Port, nil)
}
