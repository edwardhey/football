package models

/*
github.com/twitter/snowflake in golang

id =>  timestamp sequence retain businessid
          40        6         8         10

*/

import (
	"fmt"
	"sync"
	"time"
)

// var partnerId uint32

func init() {

}

const (
	nano = 1000 * 1000
)

const (
	businessTypeOrder = 1
)

const (
	TimestampBits = 40                         // timestamp
	Maxtimestamp  = -1 ^ (-1 << TimestampBits) // timestamp mask
	SequenceBits  = 6                          // sequence
	MaxSequence   = -1 ^ (-1 << SequenceBits)  // sequence mask
	RetainedBits  = 8                          // worker id
	RetainedId    = -1 ^ (-1 << RetainedBits)  // worker id mask
	BusinessBits  = 10                         // business id
	MaxBusinessId = -1 ^ (-1 << BusinessBits)  // business id mask
)

var (
	Since  int64                 = time.Date(2017, 5, 1, 0, 0, 0, 0, time.Local).UnixNano() / nano
	poolMu sync.RWMutex          = sync.RWMutex{}
	pool   map[uint64]*SnowFlake = make(map[uint64]*SnowFlake)
)

type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	businessId    uint32
	lock          sync.Mutex
}

func (sf *SnowFlake) uint64() uint64 {
	return (sf.lastTimestamp << (SequenceBits + RetainedBits + BusinessBits)) |
		(uint64(sf.sequence) << (RetainedBits + BusinessBits)) |
		// (uint64(sf.partnerId) << BusinessBits) |
		(uint64(sf.businessId))
}

func (sf *SnowFlake) Next() (uint64, error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts := timestamp()
	if ts == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & MaxSequence
		if sf.sequence == 0 {
			ts = tilNextMillis(ts)
		}
	} else {
		sf.sequence = 0
	}

	if ts < sf.lastTimestamp {
		return 0, fmt.Errorf("Invalid timestamp: %v - precedes %v", ts, sf)
	}
	sf.lastTimestamp = ts
	return sf.uint64(), nil
}

func NewSnowFlake(businessId uint32) (*SnowFlake, error) {
	if businessId < 0 || businessId > MaxBusinessId {
		return nil, fmt.Errorf("Business id %v is invalid", businessId)
	}
	return &SnowFlake{businessId: businessId}, nil
}

func timestamp() uint64 {
	return uint64(time.Now().UnixNano()/nano - Since)
}

func tilNextMillis(ts uint64) uint64 {
	i := timestamp()
	for i <= ts {
		i = timestamp()
	}
	return i
}

func GetSnowFlake(businessId uint32) (*SnowFlake, error) {
	var key uint64 = uint64(businessId) << SequenceBits
	var sf *SnowFlake
	var exist bool
	var err error
	poolMu.RLock()
	if sf, exist = pool[key]; !exist {
		poolMu.RUnlock()
		poolMu.Lock()
		// double check
		if sf, exist = pool[key]; !exist {
			sf, err = NewSnowFlake(businessId)
			if err != nil {
				poolMu.Unlock()
				return nil, err
			}
			pool[key] = sf
		}
		poolMu.Unlock()
	} else {
		poolMu.RUnlock()
	}
	return sf, err
}
