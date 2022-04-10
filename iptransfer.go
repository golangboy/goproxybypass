package main

import "sync"

type IpTransfer struct {
	sync.Mutex
	l []string
}

func (this *IpTransfer) Push(ip string) {
	for i := range this.l {
		if this.l[i] == ip {
			return
		}
	}
	this.l = append(this.l, ip)
}
func (this *IpTransfer) Exist(ip string) bool {
	for i := range this.l {
		if this.l[i] == ip {
			return true
		}
	}
	return false
}
