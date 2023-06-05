package LRUCache

import "fmt"

type LRUCache map[int]int

// the capacity allowed for inserting elements.
var capacity int

// It's used to count the elements inserted.
// If the count exceeds the capacity, it indicates that it's full
var inserted int

var lruCache LRUCache

func New(theCapacity int) (*LRUCache, error) {
	if theCapacity < 1 || theCapacity > 1000 {
		return nil, fmt.Errorf("the capacity cannot be less than 0 or higher than 1000")
	}

	capacity = theCapacity
	inserted = 0
	lruCache = make(LRUCache)

	return &lruCache, nil
}

// Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1
// Constraint: 0 <= key <= 10^3(1000)
func (l *LRUCache) Get(key int) int {
	if key < 0 || (key < 0 && key >= 1000) {
		return -1
	}
	if key > 1000 {
		return -1
	}

	return lruCache[key]
}

// Put Set or insert the value if the key is not already present. When the cache reaches its capacity,
// it should invalidate the least recently used item before inserting a new item.
// 0 <= value <= 10^5
func (l *LRUCache) Put(key int, value int) {
	if !(value < 0 || value > 0) {
		return
	}

	// if the inserted elements less than capacity to add a new element.
	if capacity > inserted {
		lruCache[key] = value
		inserted++
		return
	}

	// If the key exists in the map, it will search for the corresponding value
	// and compare it with the new value. If the new value is greater,
	// it will replace the last value for the new value.
	if l.Get(key) != -1 {
		lastValue := lruCache[key]
		newValue := value

		if newValue > lastValue {
			lruCache[key] = newValue
		}
		return
	}

	// searches the key with lower value
	var min = 100000
	var keyToReplace int
	for k, v := range lruCache {
		if v <= min {
			min = v
			keyToReplace = k
		}
	}

	// compares the lower value of the lruCache with the new value with different key
	// if the new value is less than lower value of the lruCache, it finishes
	if l.Get(keyToReplace) < value {
		return
	}

	// copies all elements of the lruCache, except the key and value that has the lower value.
	// the key with the lower value is replaced with the new key and value
	copyLruCache := make(LRUCache)
	for k, v := range lruCache {
		if k == keyToReplace {
			copyLruCache[key] = value
			continue
		}
		copyLruCache[k] = v
	}

	lruCache = copyLruCache
}

// Delete Remove the value of the key if the key exists in the cache.
// Return the value if the key exists in the cache, otherwise return -1.
func (l *LRUCache) Delete(key int) int {
	if l.Get(key) == -1 {
		return -1
	}

	deletedNumber := 0

	// copies all elements of the lruCache, except the key and value that want to delete.
	// subtracts -1 to inserted elements
	copyLruCache := make(LRUCache)
	for k, v := range lruCache {
		if k != key {
			copyLruCache[k] = v
			continue
		}
		inserted--
		deletedNumber = v
	}

	lruCache = copyLruCache

	return deletedNumber
}
