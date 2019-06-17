package etime

import "syscall"

// CurrentSecond obtain current second, use syscall for better performance
func CurrentSecond() int64 {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)
	return tv.Sec
}

// CurrentMicrosecond obtain current microsecond, use syscall for better performance
func CurrentMicrosecond() int64 {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)
	return int64(tv.Sec)*1e6 + int64(tv.Usec)
}

// CurrentMillisecond obtain current millisecond, use syscall for better performance
func CurrentMillisecond() int64 {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)
	return int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3
}
