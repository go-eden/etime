package etime

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	if NowSecond() != uint32(time.Now().Unix()) {
		t.FailNow()
	}

	ms1 := NowMillisecond()
	ms2 := time.Now().UnixNano() / 1e6
	if ms2 > ms1+1 || ms1 > ms2+1 {
		t.FailNow()
	}

	micro1 := NowMicrosecond()
	micro2 := time.Now().UnixNano() / 1e3
	if micro1 > micro2+1 || micro2 > micro1+1 {
		t.FailNow()
	}
}

// BenchmarkNowSecond2-12    	29355973	        39.9 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNowSecond2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NowSecond()
	}
}

// BenchmarkCurrentSecond-12    	30000000	        40.0 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCurrentSecond(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CurrentSecond()
	}
}

// BenchmarkNowMillisecond2-12    	27879122	        40.8 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNowMillisecond2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NowMillisecond()
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

// BenchmarkNowMicrosecond2-12    	29263656	        40.0 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNowMicrosecond2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NowMicrosecond()
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
