package snowflake

import (
	"testing"
	"time"
)

var n int64

func BenchmarkIdWorker(b *testing.B) {
	b.Run("IdWorker", benchmarkIdWorker)
}

func benchmarkIdWorker(b *testing.B) {
	var workerID int64 = 10
	lastTimestamp := time.Now().UnixMicro()
	idWorker, _ := NewIdWorker(workerID, lastTimestamp)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = idWorker.NextID()
	}
}

func TestIdWorker(t *testing.T) {
	var workerID int64 = 10
	lastTimestamp := time.Now().UnixMilli() - 1
	idWorker, _ := NewIdWorker(workerID, lastTimestamp)

	for i := 0; i < 1000; i++ {
		id, err := idWorker.NextID()
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Result: %d", id)
		}
	}
}
