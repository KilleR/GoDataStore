package datastores
//###########################
// Memory storage structures
//###########################
// Generic data store
type DataStore struct {
	sync.RWMutex
	cache map[string]string
}

func NewDataStore() *DataStore {
	return &DataStore{
		cache: make(map[string]string),
	}
}

func (ds *DataStore) set(key string, value string) {
	ds.Lock()
	defer ds.Unlock()

	ds.cache[key] = value
}

func (ds *DataStore) get(key string) string {
	ds.RLock()
	defer ds.RUnlock()

	return ds.cache[key]
}

func (ds *DataStore) setAll(data map[string]string) {
	ds.Lock()
	defer ds.Unlock()

	ds.cache = data
}

func (ds *DataStore) unset(key string) bool {
	ds.Lock()
	defer ds.Unlock()

 	_, ok := ds.cache[key]
	if ok {
		delete(ds.cache, key)
		return true
	}
	return false
}
