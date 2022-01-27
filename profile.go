package easypprof

import (
	"fmt"
	"github.com/pkg/profile"
	"testing"
)

// Profile allows you to easily run a CPU, Memory, Goroutine
// and Block profile easily
func Profile(t *testing.T, count uint64, f func()) {
	// 1 profile CPU
	p := profile.Start(profile.CPUProfile, profile.NoShutdownHook)
	for i := uint64(0); i < count; i++ {
		f()
	}
	p.Stop()

	// 2 profile MEM
	p = profile.Start(profile.MemProfile, profile.MemProfileRate(1024), profile.NoShutdownHook)
	for i := uint64(0); i < count; i++ {
		f()
	}
	p.Stop()

	// 3 profile Block
	p = profile.Start(profile.BlockProfile, profile.NoShutdownHook)
	for i := uint64(0); i < count; i++ {
		f()
	}
	p.Stop()

	// 4 profile Go routines
	p = profile.Start(profile.GoroutineProfile, profile.NoShutdownHook)
	for i := uint64(0); i < count; i++ {
		f()
	}
	p.Stop()

	// 5 profile Go heap
	p = profile.Start(profile.MemProfileHeap, profile.NoShutdownHook)
	for i := uint64(0); i < count; i++ {
		f()
	}
	p.Stop()
	report()
}

func report() {
	fmt.Println()
	fmt.Println("Profiling done")
	fmt.Println("go tool pprof -http :8080 $OUTPUT_FILE")
}
