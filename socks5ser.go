package main

import (
	"github.com/things-go/go-socks5"
	"golang.org/x/net/context"
	"golang.org/x/net/proxy"
	"net"
	"strings"
)

type Socks5Ser struct {
	up string
}

func (this *Socks5Ser) Start(listen string) {
	socks5.NewServer(
		socks5.WithDial(func(ctx context.Context, network, addr string) (net.Conn, error) {
			ip := addr[:strings.LastIndex(addr, ":")]
			byProxy := t.Exist(ip)
			if byProxy {
				dial, err := proxy.SOCKS5("tcp", this.up, nil, proxy.Direct)
				if err != nil {
					return nil, err
				}
				lg.LogStringInfo("[Proxy] " + addr)
				return dial.Dial("tcp", addr)
			}
			lg.LogStringInfo("[Direct] " + addr)
			return net.Dial(network, addr)
		}),
	).ListenAndServe("tcp", listen)
}

func (this *Socks5Ser) SetUpStream(up string) {
	this.up = up
}
