package main

import "github.com/gin-gonic/gin"

type Router struct {
	listen string
}

func (this *Router) Start() {
	g := gin.Default()
	g.GET("/add", func(context *gin.Context) {
		domain := context.Query("d")
		_ = domain
		ft.Add(domain, "*")
		lg.LogStringInfo("[Operate] Add domain: " + domain)
	})
	g.Run(this.listen)
}

func (this *Router) SetListen(listen string) {
	this.listen = listen
}
