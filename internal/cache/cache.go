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
	newLastNode := cache.last.before
	newLastNode.after = nil
	cache.last.before = nil
	cache.last = newLastNode
	delete(cache.hashmap, entry)
}

func (cache *UniversalCache) Add(entry string) {
	newEntry := UniCacheNode{}

	switch cache.entryCount {
	case 0:
		cache.first = &newEntry
		cache.last = &newEntry
		cache.hashmap[entry] = &newEntry
	case 1:
		cache.first = &newEntry
		cache.first.after = cache.last
		cache.last.before = cache.first
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

func (cache *UniversalCache) Contains(entry string) bool {
	cacheEntry, ok := cache.hashmap[entry]
	if !ok {
		cache.Add(entry)
		return false
	}
	if cacheEntry == cache.first {
		return true
	}
	if cacheEntry == cache.last {
		cache.last = cache.last.before
		cache.last.after = nil

		cacheEntry.before = nil
		cacheEntry.after = cache.first
		cache.first = cacheEntry
		return true
	}
	//reorder cacheNodes
	tempBefore := cacheEntry.before
	tempAfter := cacheEntry.after

	tempBefore.after = tempAfter
	tempAfter.before = tempBefore

	cache.first.before = cacheEntry
	cacheEntry.before = nil
	cacheEntry.after = cache.first
	cache.first = cacheEntry
	return true
}

func (cache *UniversalCache) Clear() {
	for _, node := range cache.hashmap {
		node.before = nil
		node.after = nil
	}
	cache.first = nil
	cache.last = nil
	clear(cache.hashmap)
}
