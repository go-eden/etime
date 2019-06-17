# etime

`etime` Extend golang's `time` package, to provide more and better features.

# Install

```bash
go get github.com/go-eden/etime
```
 
# `CurrentSecond`, `CurrentMicrosecond`, `CurrentMillisecond`

This feature works like Java's `System.currentTimeMillis()`, it will return `int64` value directly.

For better performance, `Current*` didn't use `time.Now()`, because it's a bit slow. 

## Demo

```go
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
```

## Performance

In my benchmark, `etime.Current*` was `40.0 ns/op`, and `time.Now()*` was `67.7 ns/op`, more details:

```
BenchmarkCurrentSecond-12         	30000000	        40.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkCurrentMillisecond-12    	50000000	        39.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkCurrentMicrosecond-12    	30000000	        39.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkNowSecond-12             	20000000	        67.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNowMillisecond-12        	20000000	        68.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkNowMicrosecond-12        	20000000	        67.8 ns/op	       0 B/op	       0 allocs/op
```

Under the same hardware environment, Java's `System.currentTimeMillis()` was like this:

```
Benchmark               Mode  Cnt   Score   Error  Units
TimestampBenchmark.now  avgt    9  25.697 ± 0.139  ns/op
```

So, there should have room for improvement in the future.

Some library may be sensitive to this `30ns` optimization, like [slf4go](https://github.com/go-eden/slf4go). 

# License

MIT