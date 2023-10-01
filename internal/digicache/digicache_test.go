package digicache

import (
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	cache := NewCache()

	key := "test_key"
	data := []byte("test_data")

	cache.Add(key, data)

	entry, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected cache to contain key %s", key)
	}

	if entry == nil {
		t.Errorf("Expected cache entry data to be %v, but got nil", data)
	}
	if string(entry) != string(data) {
		t.Errorf("Expected cache entry data to be %v, but got nil", data)
	}
}

func TestCache_Get(t *testing.T) {
	cache := NewCache()

	key := "test_key"
	data := []byte("test_data")

	cache.Add(key, data)

	entryData, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected cache to contain key %s", key)
	}

	if entryData == nil {
		t.Errorf("Expected cache entry data to be %v, but got nil", data)
	}
}
func TestCache_Reap(t *testing.T) {
	cache := NewCache()

	key1 := "test_key1"
	data1 := []byte("test_data1")

	key2 := "test_key2"
	data2 := []byte("test_data2")

	cache.Add(key1, data1)
	cache.Add(key2, data2)

	interval := time.Millisecond * 500

	// Wait for interval to pass
	time.Sleep(interval)

	cache.reap(interval)

	if _, ok := cache.cache[key1]; ok {
		t.Errorf("Expected cache to not contain key %s after reaping", key1)
	}

	if _, ok := cache.cache[key2]; ok {
		t.Errorf("Expected cache to not contain key %s after reaping", key2)
	}
}

func TestCache_ReapLoop(t *testing.T) {
	cache := NewCache()

	interval := time.Millisecond * 500

	go cache.reapLoop(interval)

	key := "test_key"
	data := []byte("test_data")

	cache.Add(key, data)

	// Wait for interval to pass
	time.Sleep(interval)

	if _, ok := cache.cache[key]; ok {
		t.Errorf("Expected cache to not contain key %s after reaping", key)
	}
}
