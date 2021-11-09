package test

import (
	"testing"
)

func BenchmarkCounter(b *testing.B) {
	b.ReportAllocs() //报告内存分配信息
	// b.N 循环次数
	for i := 0; i < b.N; i++ {
		Counter()
	}
}
func BenchmarkCounter2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Counter2()
	}
}
