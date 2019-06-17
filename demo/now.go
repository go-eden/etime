package main

import (
	"github.com/go-eden/etime"
	"time"
)

func main() {
	println(etime.CurrentSecond())
	println(etime.CurrentMillisecond())
	println(etime.CurrentMicrosecond())

	println(time.Now().Unix())
	println(time.Now().UnixNano() / 1e6)
	println(time.Now().UnixNano() / 1e3)
}
