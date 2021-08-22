package ch_gc

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestPrintStats(t *testing.T) {
	var mem runtime.MemStats
	PrintStats(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
			PrintStats(mem)
		}
	}

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
			time.Sleep(5 * time.Second)
		}
	}
	PrintStats(mem)
}
