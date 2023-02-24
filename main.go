package main

import (
	"getcahrzp.cn/router"
)

func main() {
	r := router.Router()
	r.Run(":8083") // 监听并在 0.0.0.0:8080 上启动服务
}
