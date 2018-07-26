package core

import (
	"sync"
	"time"
)

var (
	start           = time.Date(2018, 01, 01, 00, 00, 00, 0, time.UTC)
	workerBitLen    = 10
	timestampBitLen = 48
	seqBitLen       = 63 - workerBitLen - timestampBitLen

	maxSeqValue        uint64 = (1 << uint(seqBitLen)) - 1
	maxWorkerValue     uint64 = (1 << uint(workerBitLen)) - 1
	maxTimestampBitLen uint64 = (1 << uint(timestampBitLen)) - 1
)

// work|time|sequence
type IDGenerator struct {
	mutex *sync.Mutex

	worker   uint64
	sequence uint64

	lastTimestamp uint64
}

func NewIDGenerator(worker uint64) *IDGenerator {
	return &IDGenerator{
		worker: worker,
		mutex:  &sync.Mutex{},
	}
}

func (i *IDGenerator) New() uint64 {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	for {
		ts := uint64(time.Now().Sub(start) / time.Millisecond)
		if ts == i.lastTimestamp {
			if maxSeqValue == i.sequence {
				continue
			} else {
				i.sequence++
			}
		} else {
			i.sequence = 0
			i.lastTimestamp = ts
		}
		return ((i.worker & maxWorkerValue) << uint(63-workerBitLen)) | ((i.lastTimestamp & maxTimestampBitLen) << uint(seqBitLen)) | (i.sequence & maxSeqValue)
	}
}
