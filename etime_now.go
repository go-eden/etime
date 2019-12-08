package etime

import (
	"syscall"
	"time"
)

// NowSecond obtains the current second.
func NowSecond() uint32 {
	return uint32(CurrentSecond())
}

// CurrentSecond obtains the current second, use syscall for better performance
func CurrentSecond() int64 {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return time.Now().Unix()
	}
	return tv.Sec
}

// NowMillisecond obtains the current millisecond.
func NowMillisecond() int64 {
	return CurrentMillisecond()
}

// CurrentMillisecond obtains the current millisecond, use syscall for better performance
func CurrentMillisecond() int64 {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return time.Now().UnixNano() / 1e6
	}
	return tv.Sec*1e3 + int64(tv.Usec)/1e3
}

// NowMicrosecond obtains the current microsecond.
func NowMicrosecond() int64 {
	return CurrentMicrosecond()
}

// CurrentMicrosecond obtains the current microsecond, use syscall for better performance
func CurrentMicrosecond() int64 {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		return time.Now().UnixNano() / 1e3
	}
	return tv.Sec*1e6 + int64(tv.Usec)
}
