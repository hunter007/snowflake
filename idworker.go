package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	workerIDBits  int64 = 10
	sequenceBits  int64 = 12
	timestampBits int64 = 22

	baseTime    int64 = 1514736000000
	maxWorkerID int64 = 2 << workerIDBits
	minWorkerID int64 = 0
	maxSequence int64 = 2 << sequenceBits
)

type IdWorker interface {
	NextID() (int64, error)
}

func NewIdWorker(workerID, lastTimestamp int64) (IdWorker, error) {
	if workerID < minWorkerID || workerID > maxWorkerID {
		return nil, fmt.Errorf("WorkerId %d invalid", workerID)
	}

	if lastTimestamp < baseTime {
		return nil, errors.New("last timestamp too small")
	}

	return &snowflake{workerID: workerID, lastTimestamp: lastTimestamp}, nil
}

type snowflake struct {
	mutex         sync.Mutex
	sequence      int64
	workerID      int64
	lastTimestamp int64
}

func (sf *snowflake) NextID() (int64, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	nowTime := time.Now().UnixMilli()
	if nowTime < sf.lastTimestamp {
		return 0, fmt.Errorf("clock is moving backwards. Rejecting requests until %d", sf.lastTimestamp)
	}

	if nowTime == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & maxSequence
		if sf.sequence == 0 {
			for nowTime <= sf.lastTimestamp {
				nowTime = time.Now().UnixMilli()
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTimestamp = nowTime
	return (nowTime-baseTime)<<timestampBits | sf.workerID<<sequenceBits | sf.sequence, nil
}
