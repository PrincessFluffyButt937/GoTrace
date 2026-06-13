package cache

type Cache struct {
	SN *UniversalCache
	HU *UniversalCache
}

type UniversalCache struct {
	hashmap    map[string]*UniCacheNode
	casheSize  int
	entryCount int
	first      *UniCacheNode
	last       *UniCacheNode
}

type UniCacheNode struct {
	entry  string
	before *UniCacheNode
	after  *UniCacheNode
}

func InitCache(SNsize, HUsize int) *Cache {
	SNcache := UniversalCache{
		hashmap:    make(map[string]*UniCacheNode),
		casheSize:  SNsize,
		entryCount: 0,
	}
	HUcache := UniversalCache{
		hashmap:    make(map[string]*UniCacheNode),
		casheSize:  SNsize,
		entryCount: 0,
	}

	finalCache := Cache{
		SN: &SNcache,
		HU: &HUcache,
	}
	return &finalCache
}

func (cache *UniversalCache) DelLastNode(entry string) {
	cache.last = cache.last.before
	cache.last.after = nil
	delete(cache.hashmap, entry)
}

func (cache *UniversalCache) Add(entry string) {
	newEntry := UniCacheNode{
		entry: entry,
	}
	switch cache.entryCount {
	case 0:
		cache.first = &newEntry
		cache.last = &newEntry
		cache.hashmap[entry] = &newEntry
	case 1:
		cache.first = &newEntry
		newEntry.after = cache.last
		cache.last.before = &newEntry
		cache.hashmap[entry] = &newEntry
	default:
		newEntry.after = cache.first
		cache.first.before = &newEntry
		cache.first = &newEntry
		cache.hashmap[entry] = &newEntry
		if cache.casheSize == cache.entryCount {
			cache.DelLastNode(entry)
			cache.entryCount--
		}
	}
	cache.entryCount++
}
