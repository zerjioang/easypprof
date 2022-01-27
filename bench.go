package easypprof

import "testing"

// Bench allows you to easily run benchmark adding basic
// boilerplate code
func Bench(b *testing.B, f func()) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f()
	}
}
