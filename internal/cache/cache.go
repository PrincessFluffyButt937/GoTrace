package cache

type Cache struct {
	SN   *UniversalCache
	HU   *UniversalCache
	File *FileCache
}

// map [SN][]Filepath

//UniversalCache

type UniversalCache struct {
	hashmap    map[string]*UniCacheNode
	casheSize  int
	entryCount int
	first      *UniCacheNode
	last       *UniCacheNode
}

type UniCacheNode struct {
	key    string
	before *UniCacheNode
	after  *UniCacheNode
}

func InitCache(SNsize, HUsize, FileSize int) *Cache {
	SNcache := UniversalCache{
		hashmap:    make(map[string]*UniCacheNode),
		casheSize:  SNsize,
		entryCount: 0,
	}
	HUcache := UniversalCache{
		hashmap:    make(map[string]*UniCacheNode),
		casheSize:  HUsize,
		entryCount: 0,
	}
	FileNameCache := FileCache{
		hashmap:    make(map[string]*FileCacheNode),
		casheSize:  FileSize,
		entryCount: 0,
	}

	finalCache := Cache{
		SN:   &SNcache,
		HU:   &HUcache,
		File: &FileNameCache,
	}
	return &finalCache
}

//Universal Cache

func (cache *UniversalCache) DelLastNode(entry string) {
	newLastNode := cache.last.before
	newLastNode.after = nil
	cache.last.before = nil
	cache.last = newLastNode
	delete(cache.hashmap, entry)
}

func (cache *UniversalCache) Add(entry string) {
	newEntry := UniCacheNode{
		key: entry,
	}

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
			cache.DelLastNode(cache.last.key)
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
	cache.entryCount = 0
	clear(cache.hashmap)
}

//File cache

type FileCache struct {
	hashmap    map[string]*FileCacheNode
	casheSize  int
	entryCount int
	first      *FileCacheNode
	last       *FileCacheNode
}

type FileCacheNode struct {
	key    string
	before *FileCacheNode
	after  *FileCacheNode
	files  map[string]struct{}
}

func (cache *FileCache) DelLastNode(entry string) {
	newLastNode := cache.last.before
	newLastNode.after = nil
	cache.last.before = nil
	clear(cache.last.files)
	cache.last = newLastNode
	delete(cache.hashmap, entry)
}

func (cache *FileCache) Add(entry, fileName string) {
	newEntry := FileCacheNode{
		key:   entry,
		files: make(map[string]struct{}),
	}
	newEntry.files[fileName] = struct{}{}

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
			cache.DelLastNode(cache.last.key)
			cache.entryCount--
		}
	}
	cache.entryCount++
}

func (cache *FileCache) Contains(entry, fileName string) bool {
	cacheEntry, ok := cache.hashmap[entry]
	if !ok {
		cache.Add(entry, fileName)
		return false
	}

	//Write fileName into cache
	cacheEntry.files[fileName] = struct{}{}

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

func (cache *FileCache) Fetch(entry string) []string {
	cashedEntries, ok := cache.hashmap[entry]
	if !ok {
		return nil
	}

	finalSlice := make([]string, len(cashedEntries.files))

	for filePath, _ := range cashedEntries.files {
		finalSlice = append(finalSlice, filePath)
	}

	//delete cache entry method??

	return finalSlice
}

func (cache *FileCache) Clear() {
	for _, node := range cache.hashmap {
		node.before = nil
		node.after = nil
		clear(node.files)
	}
	cache.first = nil
	cache.last = nil
	cache.entryCount = 0
	clear(cache.hashmap)
}
