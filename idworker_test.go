package snowflake

import (
	"testing"
	"time"
)

func BenchmarkIdWorker(b *testing.B) {
	b.Run("IdWorker-10", benchmark10IdWorker)
	b.Run("IdWorker-100", benchmark100IdWorker)
}

func benchmark10IdWorker(b *testing.B) {
	var workerID int64 = 10
	lastTimestamp := time.Now().UnixMicro()
	idWorker, _ := NewIdWorker(workerID, lastTimestamp)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = idWorker.NextID()
	}
}

func benchmark100IdWorker(b *testing.B) {
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

	for i := 0; i < 100; i++ {
		id, err := idWorker.NextID()
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Result: %d", id)
		}
	}
}
