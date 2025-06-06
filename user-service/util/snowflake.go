package util

import (
	"errors"
	"sync"
	"time"
)

const (
	workerIDBits     = 5
	datacenterIDBits = 5
	sequenceBits     = 12

	maxWorkerID     = -1 ^ (-1 << workerIDBits)
	maxDatacenterID = -1 ^ (-1 << datacenterIDBits)
	maxSequence     = -1 ^ (-1 << sequenceBits)

	workerIDShift      = sequenceBits
	datacenterIDShift  = sequenceBits + workerIDBits
	timestampLeftShift = sequenceBits + workerIDBits + datacenterIDBits

	// 2020-01-01 00:00:00 UTC 作为时间戳起点
	epoch = int64(1577836800000)
)

// Snowflake 雪花算法结构体
type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	datacenterID  int64
	sequence      int64
}

// NewSnowflake 创建一个新的 Snowflake 实例
func NewSnowflake(workerID, datacenterID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker GoodID 超出范围")
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, errors.New("数据中心 GoodID 超出范围")
	}
	return &Snowflake{
		lastTimestamp: -1,
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
	}, nil
}

// NextID 生成下一个唯一 GoodID
func (s *Snowflake) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixNano() / 1e6
	if timestamp < s.lastTimestamp {
		return 0, errors.New("时钟回拨，拒绝生成 GoodID")
	}
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// 当前毫秒内序列号用完，等待下一毫秒
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampLeftShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id, nil
}
