package main

import (
	"github.com/miekg/dns"
)

type DnsSer struct {
	up string
}

func (this *DnsSer) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	c := new(dns.Client)
	useRule := false
	for i := range r.Question {
		lg.LogStringInfo("[DNS Query] " + r.Question[i].String())
		if ft.IsBypass(r.Question[i].Name) {
			useRule = true
			break
		}
	}
	var rr *dns.Msg
	var err error
	if useRule {
		rr, _, err = c.Exchange(r, this.up)
		if err != nil {
			return
		}
		for i := range rr.Answer {
			if v, ok := rr.Answer[i].(*dns.A); ok {
				lg.LogStringInfo("[DNS]" + v.A.String())
				t.Push(v.A.String())
			} else if v, ok := rr.Answer[i].(*dns.AAAA); ok {
				lg.LogStringInfo("[DNS]" + v.AAAA.String())
				t.Push(v.AAAA.String())
			}
		}
	} else {
		rr, _, err = c.Exchange(r, *directDnsUpStream)
		if err != nil {
			return
		}
	}
	w.WriteMsg(rr)
}

func (this *DnsSer) Start() {
	s := dns.Server{
		Addr: ":53",
		Net:  "udp",
	}
	s.Handler = this
	s.ListenAndServe()
}
func (this *DnsSer) SetUpStream(up string) {
	this.up = up
}
