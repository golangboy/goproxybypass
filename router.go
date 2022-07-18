package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Router struct {
	listen string
}

func (this *Router) Start() {
	g := gin.Default()
	g.GET("/add", func(context *gin.Context) {
		domain := context.Query("d")
		ipv4 := context.Query("v4")
		if ipv4 != "" {
			t.Push(ipv4)
			lg.LogStringInfo("[Operate] Add ipv4: " + ipv4)
			return
		}
		_ = domain
		ft.Add(domain, "*")
		lg.LogStringInfo("[Operate] Add domain: " + domain)
	})
	g.GET("/switch", func(context *gin.Context) {
		globalProxy = !globalProxy
		context.Writer.WriteString("global proxy:" + strconv.FormatBool(globalProxy))
	})

	g.Run(this.listen)
}

func (this *Router) SetListen(listen string) {
	this.listen = listen
}
