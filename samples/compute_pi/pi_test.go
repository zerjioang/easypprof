package main

import (
	"github.com/zerjioang/easypprof"
	"testing"
)

func TestPi(t *testing.T) {
	easypprof.Profile(t, 5000000, func() {
		computePi(200)
	})
}

func BenchmarkPi(b *testing.B) {
	easypprof.Bench(b, func() {
		computePi(200)
	})
}
