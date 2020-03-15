package main

import (
	"dir/route"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var path string

func init() {
	flag.StringVar(&path, "p", "/", "log in user")
	flag.Parse()
}

// go run main.go -p /Users/admin/www/learn/go/src/dir/
func main() {
	//path = "/Users/admin/www/learn/go/src/dir/"
	server := &http.Server{
		Handler: route.MyHandler{
			RootPath: path,
		}, // 使用实现 http.Handler 的结构处理 HTTP 数据
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("start http server fail:", err)
	}
}
