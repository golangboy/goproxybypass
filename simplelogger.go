package main

import (
	"log"
)

type SimpleLogger struct {
}

func (this *SimpleLogger) LogStringInfo(str string) {
	log.Default().Println(str)
}

func (this *SimpleLogger) LogStringDebug(str string) {

}

func (this *SimpleLogger) LogStringError(str string) {

}

func (this *SimpleLogger) LogStringWarning(str string) {

}
