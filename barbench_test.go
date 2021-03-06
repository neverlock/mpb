package mpb

import (
	"io/ioutil"
	"testing"
)

func benchmarkSingleBar(total int) {
	p := New(Output(ioutil.Discard))
	bar := p.AddBar(int64(total))
	for i := 0; i < total; i++ {
		bar.Increment()
	}
	p.Stop()
}

func BenchmarkSingleBar100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkSingleBar(100)
	}
}

func BenchmarkSingleBar1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkSingleBar(1000)
	}
}

func BenchmarkSingleBar10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkSingleBar(10000)
	}
}

func BenchmarkIncrSingleBar(b *testing.B) {
	p := New(Output(ioutil.Discard))
	bar := p.AddBar(int64(b.N))
	for i := 0; i < b.N; i++ {
		bar.Increment()
	}
	p.Stop()
}
