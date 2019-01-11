package logs

import (
	"log"
	"os"
)

var (
	Info 			*log.Logger
	Error		*log.Logger
)

/*
	初始化日志记录器
 */
func Init() {
	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stdout,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info.Println("日志记录器初始化完成")
}
