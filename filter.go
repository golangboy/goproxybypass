package main

import "strings"

type SimpleFilter struct {
	domainList []string
}

func (this *SimpleFilter) IsBypass(domain string) bool {
	for _, v := range this.domainList {
		if strings.Index(domain, v) != -1 {
			return true
		}
	}
	return false
}
func (this *SimpleFilter) Add(domain string, rule string) {
	for _, v := range this.domainList {
		if v == domain {
			return
		}
	}
	this.domainList = append(this.domainList, domain)
}
func (this *SimpleFilter) Del(rule string) {
	//for i, v := range this.domainList {
	//	if v == rule {
	//		this.domainList = append(this.domainList[:i], this.domainList[i+1:]...)
	//		return
	//	}
	//}
}
