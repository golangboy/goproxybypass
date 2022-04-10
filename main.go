package main

import "flag"

var lg LoggerInf
var t IpTransfer
var ft SimpleFilter
var dnsUpStream *string
var socks5UpSteam *string
var webAddr *string
var socks5Addr *string

func main() {
	d := DnsSer{}
	r := Router{}
	s := Socks5Ser{}
	lg = &SimpleLogger{}
	dnsUpStream = flag.String("d", "8.8.8.8:53", "trust dns upstream")
	socks5UpSteam = flag.String("s", "127.0.0.1:1080", "trust socks5 upstream")
	webAddr = flag.String("w", ":80", "web bind address")
	socks5Addr = flag.String("k", ":7777", "socks5 bind address")
	flag.Parse()
	d.SetUpStream(*dnsUpStream)
	r.SetListen(*webAddr)
	s.SetUpStream(*socks5UpSteam)
	go s.Start(*socks5Addr)
	go d.Start()
	r.Start()
}