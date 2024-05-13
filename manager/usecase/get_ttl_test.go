package usecase

import (
	"fmt"
	"github.com/saaitt/gomem/manager/domain"
	"runtime"
	"testing"
	"time"
)

func BenchmarkGetTTl(b *testing.B) {
	rt := new(domain.TTLRadixTree)
	k := "a"
	val := "sadfasd"
	ttl := time.Duration(10)
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	for i := 0; i < b.N; i++ {
		rt.Insert(&k, &val, &ttl)
	}
	runtime.ReadMemStats(&m2)
	fmt.Println()
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)

}
