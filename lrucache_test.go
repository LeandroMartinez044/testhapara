package LRUCache

import (
	"testing"
)

func TestNew_WhenCapacityIsLessThan1orHigherThan1000(t *testing.T) {
	capacity := 0
	lruCache, err := New(capacity)
	if lruCache != nil || err == nil {
		t.Fatalf(`New(0) = %v, %v, error`, lruCache, err)
	}
}

func TestNew_OK(t *testing.T) {
	capacity := 4
	lruCache, _ := New(capacity)
	lruCache.Put(15, 9)
}

func TestGet_OK(t *testing.T) {
	capacity := 4
	lruCache, _ := New(capacity)
	lruCache.Put(1, 10)

	num := lruCache.Get(1)

	if num != 10 {
		t.Fatalf(`Get(1) = -1, want 10`)
	}
}

func TestGet_WhenKeyIsLessThan0orHigherThan1000(t *testing.T) {
	capacity := 4
	lruCache, _ := New(capacity)
	lruCache.Put(1, 10)

	val1 := lruCache.Get(1001)

	if val1 != -1 {
		t.Fatalf(`Get(1001) = %d, want -1`, val1)
	}

	val2 := lruCache.Get(-1)

	if val2 != -1 {
		t.Fatalf(`Get(-1) = %d, want -1`, val2)
	}
}

func TestPut_OK(t *testing.T) {
	capacity := 4
	lruCache, _ := New(capacity)
	lruCache.Put(1, 10)
	lruCache.Put(15, 11)
	lruCache.Put(7, 9)
	lruCache.Put(3, 12)
	lruCache.Put(5, 13)
	lruCache.Put(3, 13)

	if lruCache.Get(7) == -1 || inserted != 4 {
		t.Fatalf(`Get(7) = %q, want "-1"`, lruCache.Get(7))
	}

}

func TestDelete_OK(t *testing.T) {
	capacity := 4
	lruCache, _ := New(capacity)
	lruCache.Put(0, 9)
	lruCache.Put(1, 10)
	lruCache.Put(2, 11)
	lruCache.Put(3, 12)

	val1 := lruCache.Delete(1)
	if val1 != 10 || inserted != 3 {
		t.Fatalf(`Delete(1) = %q, want "10"`, val1)
	}

	val2 := lruCache.Delete(2)
	if val2 != 11 || inserted != 2 {
		t.Fatalf(`Delete(2) = %q, want "11"`, val2)
	}

	lruCache.Put(10, 15)
	lruCache.Put(11, 16)
	lruCache.Put(12, 18)
	if inserted != 4 {
		t.Fatalf(`Delete(1) = %q, want "11"`, val2)
	}
}
