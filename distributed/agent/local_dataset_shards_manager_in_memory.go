package agent

import (
	"sync"
	"time"

	"github.com/chrislusf/gleam/util"
)

type trackedChannel struct {
	incomingChannel  *util.Piper
	outgoingChannels []*util.Piper
	index            int
	wg               *sync.WaitGroup
	lastWriteAt      time.Time
	isClosed         bool
}

func newTrackedChannel(readerCount int) *trackedChannel { _ = "STUB: not implemented"; return nil }

func (tc *trackedChannel) borrowChannel() *util.Piper { _ = "STUB: not implemented"; return nil }

type LocalDatasetShardsManagerInMemory struct {
	sync.Mutex
	name2Channel     map[string]*trackedChannel
	name2ChannelCond *sync.Cond
}

func NewLocalDatasetShardsManagerInMemory() *LocalDatasetShardsManagerInMemory {
	_ = "STUB: not implemented"
	return nil
}

func (m *LocalDatasetShardsManagerInMemory) doDelete(name string) {
	_ = "STUB: not implemented"
	return
}

func (m *LocalDatasetShardsManagerInMemory) CreateNamedDatasetShard(name string, readerCount int) *trackedChannel {
	_ = "STUB: not implemented"
	return nil
}

// println("setting", name, "to", m, m.name2Channel[name])

func (m *LocalDatasetShardsManagerInMemory) WaitForNamedDatasetShard(name string) *util.Piper {
	_ = "STUB: not implemented"
	return nil
}

// println("found existing channel", name, "closed:", tc.isClosed)

// println("waiting for", name, m, m.name2Channel[name])

// println("woke up for", name, m, m.name2Channel[name])

func (m *LocalDatasetShardsManagerInMemory) Cleanup(name string) { _ = "STUB: not implemented"; return }

// purge executor status older than 24 hours to save memory
func (m *LocalDatasetShardsManagerInMemory) purgeExpiredEntries() {
	_ = "STUB: not implemented"
	return
}
