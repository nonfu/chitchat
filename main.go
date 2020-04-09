package main

import (
    . "github.com/xueyuanjun/chitchat/config"
    . "github.com/xueyuanjun/chitchat/routes"
    "log"
    "net/http"
)

func main()  {
    startWebServer()
}

// 通过指定端口启动 Web 服务器
func startWebServer()  {
    config := LoadConfig()
    r := NewRouter() // 通过 router.go 中定义的路由器来分发请求

    // 处理静态资源文件
    assets := http.FileServer(http.Dir(config.App.Static))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

    http.Handle("/", r)

    log.Println("Starting HTTP service at " + config.App.Address)
    err := http.ListenAndServe(config.App.Address, nil)

    if err != nil {
        log.Println("An error occured starting HTTP listener at " + config.App.Address)
        log.Println("Error: " + err.Error())
    }
}
