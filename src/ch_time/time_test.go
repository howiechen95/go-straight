package ch_time

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	for {
		func() {
			flush := false
			timer := time.NewTimer(1 * time.Second)
			for !flush {
				select {
				//case v, run := <-storeCh:
				//	if !run {
				//		return
				//	}
				//	if _, ok := m[v.Key]; !ok {
				//		m[v.Key] = make(map[uint64]bool, 100)
				//	}
				//	m[v.Key][v.UID] = true
				case <-timer.C:
					flush = true
				}
			}
			t.Log(time.Now().String())
		}()
	}

}
