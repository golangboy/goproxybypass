package main

import (
	"github.com/miekg/dns"
)

type DnsSerInf interface {
	Start()
	ServeDNS(w dns.ResponseWriter, r *dns.Msg)
	SetUpStream(up string)
}
type RouterInf interface {
	Start()
	SetListen(listen string)
}

type IpTransferInf interface {
	Push(ip string)
	Exist(ip string) bool
}

type FilterInf interface {
	IsBypass(domain string) bool
	Add(domain string, rule string)
	Del(rule string)
}

type Socks5TransferInf interface {
	Start(listen string)
	SetUpStream(up string)
}

type LoggerInf interface {
	LogStringInfo(str string)
	LogStringError(str string)
	LogStringDebug(str string)
	LogStringWarning(str string)
}
