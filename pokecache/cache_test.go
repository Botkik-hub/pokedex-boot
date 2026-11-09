package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(10)

	val, found := cache.Get("test")
	if found {
		t.Errorf("found a key in empty cache, val: %s", string(val))
	}
	testStr := "test1"
	testKey := "key1"
	cache.Add(testKey, []byte(testStr))

	val, found = cache.Get(testKey)

	if !found || string(val) != testStr {
		t.Errorf("not found a key in cache or vals did not match; expected: %s, actual: %s", testStr, string(val))
	}

	testStr2 := "test2"
	testKey2 := "key2"
	cache.Add(testKey2, []byte(testStr2))

	val, found = cache.Get(testKey)
	if !found || string(val) != testStr {
		t.Errorf("not found a key1 in cache or vals did not match with two keys; expected: %s, actual: %s", testStr, string(val))
	}

	val, found = cache.Get(testKey2)
	if !found || string(val) != testStr2 {
		t.Errorf("not found a key2 in cache or vals did not match with two keys; expected: %s, actual: %s", testStr2, string(val))
	}

	testStr = "newTest1"

	cache.Add(testKey, []byte(testStr))
	val, found = cache.Get(testKey)
	if !found || string(val) != testStr {
		t.Errorf("not found a key1 in cache or vals did not match with two keys, val1 replaced; expected: %s, actual: %s", testStr, string(val))
	}
}

func TestCacheReap(t *testing.T) {
	cache := NewCache(1)

	val, found := cache.Get("test")
	if found {
		t.Errorf("found a key in empty cache, val: %s", string(val))
	}
	testStr := "test1"
	testKey := "key1"
	cache.Add(testKey, []byte(testStr))

	val, found = cache.Get(testKey)

	if !found || string(val) != testStr {
		t.Errorf("not found a key in cache or vals did not match; expected: %s, actual: %s", testStr, string(val))
	}
	time.Sleep(2 * time.Second)
	val, found = cache.Get(testKey)

	if found {
		t.Errorf("found a key when it was suppopse to reap in cahce: %s", string(val))
	}
}
