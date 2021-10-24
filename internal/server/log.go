package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

// Append return commit logs offset（配列の初期位置からの場所）
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock() // 遅延
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read return commit record
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock() // 遅延
	// 初期位置から数えて存在しない場合はからのレコードの例外を返す
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value []byte `json:"value"` // base64でbyteを表すものをPOSTする必要がある ex. TGV0J3MgR28gIzI
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")