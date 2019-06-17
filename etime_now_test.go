package etime

import (
	"testing"
	"time"
)

// BenchmarkCurrentSecond-12    	30000000	        40.0 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCurrentSecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CurrentSecond()
	}
}

// BenchmarkCurrentMillisecond-12    	30000000	        40.5 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCurrentMillisecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CurrentMillisecond()
	}
}

// BenchmarkCurrentMicrosecond-12    	30000000	        40.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCurrentMicrosecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CurrentMicrosecond()
	}
}

// BenchmarkNowSecond-12    	20000000	        67.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNowSecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Now().Unix()
	}
}

// BenchmarkNowMillisecond-12    	20000000	        67.7 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNowMillisecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Now().UnixNano() / 1e6
	}
}

// BenchmarkNowMicrosecond-12    	20000000	        67.6 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNowMicrosecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Now().Unix() / 1e3
	}
}
