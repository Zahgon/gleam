package agent

import (
	"sync"

	"github.com/chrislusf/gleam/distributed/store"
)

type LocalDatasetShardsManager struct {
	sync.Mutex
	dir            string
	port           int
	name2Store     map[string]store.DataStore
	name2StoreCond *sync.Cond
}

func NewLocalDatasetShardsManager(dir string, port int) *LocalDatasetShardsManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *LocalDatasetShardsManager) doDelete(name string) {
	_ = "STUB: not implemented"

	// println("deleting from LocalDatasetShardsManager:", name)
	return
}

func (m *LocalDatasetShardsManager) DeleteNamedDatasetShard(name string) {
	_ = "STUB: not implemented"

	// println("locking LocalDatasetShardsManager to delete", name)
	return
}

// println("locked LocalDatasetShardsManager to delete", name)

func (m *LocalDatasetShardsManager) CreateNamedDatasetShard(name string) store.DataStore {
	_ = "STUB: not implemented"
	return *new(store.DataStore)
}

// println(name, "is broadcasting...")

func (m *LocalDatasetShardsManager) WaitForNamedDatasetShard(name string) store.DataStore {
	_ = "STUB: not implemented"
	return *new(store.DataStore)
}

// println(name, "is waiting to read...")

// purge executor status older than 24 hours to save memory
func (m *LocalDatasetShardsManager) purgeExpiredEntries() { _ = "STUB: not implemented"; return }
